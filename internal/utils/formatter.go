package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/term"
)

func FormatJSON(entry map[string]interface{}, pretty bool) ([]byte, error) { // todo: agregar colores q diferencien a los campos de su contenido
	if pretty {
		return json.MarshalIndent(entry, "", "  ")
	}
	return json.Marshal(entry)
}

func FormatText(entry map[string]interface{}, level string, pretty bool) string {
	var sb strings.Builder

	//timestamp
	ts, _ := entry["ts"].(string)
	if ts == "" {
		ts = time.Now().Format(time.RFC3339)
	}
	sb.WriteString(ts)

	sb.WriteString(" ")
	if isTerminal() {
		sb.WriteString(applyColor(level, fmt.Sprintf(" [%s]", level)))
	} else {
		sb.WriteString(fmt.Sprintf("[%s]", level))
	}

	// layer
	if layer, ok := entry["layer"].(string); ok && layer != "" {
		sb.WriteString(fmt.Sprintf(" [%s]", strings.ToUpper(layer)))
	} else {
		sb.WriteString(" [UNKNOWN]")
	}

	// message
	if msg, ok := entry["msg"].(string); ok && msg != "" {
		sb.WriteString(fmt.Sprintf(" %s", msg))
	}

	// Caller
	if caller, ok := entry["caller"].(string); ok && caller != "" {
		sb.WriteString(fmt.Sprintf(" (%s)", filepath.Base(caller)))
	}

	for k, v := range entry {
		if k == "ts" || k == "lvl" || k == "msg" || k == "layer" || k == `caller` {
			continue
		}
		sb.WriteString(fmt.Sprintf(" %s: %v", k, v))
	}

	return sb.String()
}

func isTerminal() bool {
	fd := int(os.Stdout.Fd())
	return term.IsTerminal(fd)
}

// Aplica color al nivel de log (solo si es TTY)
func applyColor(level, text string) string {
	switch level {
	case "DEBUG":
		return fmt.Sprintf("\x1b[36m%s\x1b[0m", text)
	case "INFO":
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", text)
	case "WARN":
		return fmt.Sprintf("\x1b[33m%s\x1b[0m", text)
	case "ERROR":
		return fmt.Sprintf("\x1b[31m%s\x1b[0m", text)
	case "FATAL":
		return fmt.Sprintf("\x1b[41m\x1b[37m%s\x1b[0m", text)
	default:
		return text
	}
}

//todo: Mejorar la serialización de campos complejos:

func FormatField(field any, pretty bool) string {
	// Si es una estructura o un mapa, lo convertimos a JSON
	j, err := json.Marshal(field)
	if err != nil {
		return fmt.Sprintf("%v", field)
	}
	if pretty {
		var indented bytes.Buffer
		json.Indent(&indented, j, "", "  ")
		return indented.String()
	}
	return string(j)
}

// Formatea un log completo (timestamp, nivel, mensaje, y otros campos)
func FormatLog(entry map[string]any, level string, pretty bool) string {
	if entry["lvl"] == "JSON" {
		// Si el formato es JSON, utilizamos el formato JSON
		j, _ := FormatJSON(entry, pretty)
		return string(j)
	}
	// Si no es JSON, usamos el formato de texto
	return FormatText(entry, level, pretty)
}
