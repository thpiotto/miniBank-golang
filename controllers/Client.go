package controllers

import "fmt"

type Client struct {
	Name          string
	LastName      string
	CPF           string
	Ocupation     string
	Role          string
}

func (cli *Client) GeneralInfo() {
	fmt.Println("[Name]:", cli.Name)
	fmt.Println("[LastName]:", cli.LastName)
	fmt.Println("[CPF]:", cli.CPF)
	fmt.Print("[Ocupation]:", cli.Ocupation)
	fmt.Print("[Role]:", cli.Role)
}

func (cli *Client) CompletedName() string {
	return cli.Name + " " + cli.LastName
}
