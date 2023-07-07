package ux

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
)

type PromptContentString struct {
	Label           string
	Default         *string
	RegexValidation *regexp.Regexp
	Sensitive       bool
}

type PromptContentSelect struct {
	Label        string
	DefaultIndex int
	Options      []string
}

func PromptGetString(pc PromptContentString) string {
	if pc.RegexValidation == nil {
		pc.RegexValidation = regexp.MustCompile(`^.+$`)
	}

	if pc.Default == nil {
		pc.Default = new(string)
	}

	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("Field must not be empty")
		}
		if !pc.RegexValidation.MatchString(input) {
			return errors.New(fmt.Sprintf("The field value does not match the required format: %s", pc.RegexValidation.String()))
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	var prompt promptui.Prompt
	if pc.Sensitive {
		prompt = promptui.Prompt{
			Label:     pc.Label,
			Templates: templates,
			Validate:  validate,
			Default:   *pc.Default,
			Mask:      '*',
		}
	} else {
		prompt = promptui.Prompt{
			Label:     pc.Label,
			Templates: templates,
			Validate:  validate,
			Default:   *pc.Default,
		}
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func PromptGetSelect(pc PromptContentSelect) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label:        pc.Label,
			Items:        pc.Options,
			HideSelected: false,
			CursorPos:    pc.DefaultIndex,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
