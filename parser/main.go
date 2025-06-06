package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type Document struct {
	Text string `json:"text"`
}

// Struct para armazenar tudo que você quer extrair
type ParsedExam struct {
	PatientName    string
	Age            int
	DateCollected  time.Time
	ExamName       string
	ResultValue    string
	ResultUnit     string
	Method         string
	ReferenceRange string
	// estenda com o que precisar…
}

// Padrões regex para diferentes campos
var (
	rePatient   = regexp.MustCompile(`Paciente:\s*(.+)`)
	reAge       = regexp.MustCompile(`Idade:\s*(\d+)`)
	reCollected = regexp.MustCompile(`Coletado:\s*([\d/]+\s*[\d:]+)`)
	reMethod    = regexp.MustCompile(`Método:\s*(.+)`)
	reRange     = regexp.MustCompile(`Valor de referência\s*(.+)`) // captura tudo até a próxima linha
	// vamos capturar o primeiro exame e seu valor:
	reExamName = regexp.MustCompile(`^([A-Z0-9À-Ú ]+):\s*$`)
	reResult   = regexp.MustCompile(`^([\d.,]+)\s*([^\s]+)`)
)

func parseDocument(text string) (*ParsedExam, error) {
	p := &ParsedExam{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	var lastExamMatch string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if m := rePatient.FindStringSubmatch(line); len(m) > 1 {
			p.PatientName = m[1]
		}
		if m := reAge.FindStringSubmatch(line); len(m) > 1 {
			fmt.Sscanf(m[1], "%d", &p.Age)
		}
		if m := reCollected.FindStringSubmatch(line); len(m) > 1 {
			// assume layout "02/05/2025 09:46"
			t, _ := time.Parse("02/01/2006 15:04", m[1])
			p.DateCollected = t
		}
		if m := reMethod.FindStringSubmatch(line); len(m) > 1 {
			p.Method = m[1]
		}
		if m := reRange.FindStringSubmatch(line); len(m) > 1 {
			p.ReferenceRange = m[1]
		}
		// captura o nome do exame (linha em caps + dois-pontos)
		if m := reExamName.FindStringSubmatch(line); len(m) > 1 {
			lastExamMatch = m[1]  // guarda para associar ao próximo valor
			if p.ExamName == "" { // só o primeiro, ou adaptação para vários
				p.ExamName = m[1]
			}
		}
		// logo em seguida vem o valor+unidade
		if lastExamMatch != "" && p.ResultValue == "" {
			if m := reResult.FindStringSubmatch(line); len(m) > 1 {
				p.ResultValue = m[1]
				p.ResultUnit = m[2]
				lastExamMatch = "" // zera pra não pegar outras linhas
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return p, nil
}

func main() {
	data, _ := os.ReadFile("document.json")
	var doc Document
	_ = json.Unmarshal(data, &doc)

	parsed, err := parseDocument(doc.Text)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Printf("Patient: %s\nAge: %d\nCollected: %s\n",
		parsed.PatientName, parsed.Age, parsed.DateCollected)
	fmt.Printf("Exam: %s → %s %s\nMethod: %s\nRefRange: %s\n",
		parsed.ExamName, parsed.ResultValue, parsed.ResultUnit,
		parsed.Method, parsed.ReferenceRange)
}
