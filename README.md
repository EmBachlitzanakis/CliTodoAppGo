# CLI Todo List with Priority Suggestion using LangChain and Ollama 2

This is a command-line interface (CLI) Todo List application that not only helps you manage your tasks but also uses AI to suggest the priority level of each task. The app uses [LangChain](https://github.com/hwchase17/langchain) and the [Ollama 2 model](https://ollama.com/) to suggest priority levels (high, medium, low) for each task based on the provided task description.

## Features

- **Add, edit, toggle, and delete tasks**: Manage your tasks easily via the CLI.
- **AI-Powered Task Prioritization**: Get suggested priority levels for your tasks using LangChain and Ollama 2.
- **Categorize tasks**: Assign your own priority level or use the AI's suggestion.
- **Task List Management**: View all tasks sorted by priority, due date, or status.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

To set up the CLI Todo List application with AI-Powered Task Prioritization, follow these steps:

### Prerequisites

Before you get started, ensure you have the following installed on your machine:

- **Ollama**: [Download and install Ollama](https://ollama.com/download).
- **Go**: [Download and install Go](https://golang.org/dl/).

## Getting Started

Follow these steps to set up the project and run the example:

### Step 1: Initialize Ollama

Open your terminal and execute the following command:


$ ollama run llama2

- go run github.com/tmc/langchaingo/examples/ollama-completion-example@main

You should receive an output similar to this:

- The first human to set foot on the moon was Neil Armstrong, an American astronaut, who stepped onto the lunar surface.

those are the instactions from :

[LangChainGo](https://tmc.github.io/langchaingo/docs/getting-started/guide-ollama)

### Step 2: Clone repository 

gh repo clone EmBachlitzanakis/CliTodoAppGo

## Usage  

Below I will show some examples of how to check your list : 
![image_2024_09_30T11_39_15_067Z](https://github.com/user-attachments/assets/b12c9860-3981-4df5-8b8a-a8f21729dfb2)

Mark a task as done :
![image_2024_09_30T11_40_06_775Z](https://github.com/user-attachments/assets/b774876f-b580-4fb3-86d1-46302cafecc3)

Delete a task:
![image_2024_09_30T11_40_56_248Z](https://github.com/user-attachments/assets/f57ffd0d-c80e-430c-98ed-66054a468d01)

Add a task:
![image_2024_09_30T11_41_31_785Z](https://github.com/user-attachments/assets/2b94643a-50f1-432b-af83-b307d8325802)

Ask AI to determine the priority level of your tasks : 
![image_2024_09_30T11_42_40_841Z](https://github.com/user-attachments/assets/9c0bbec6-10af-4049-ac23-f3a57986f31a)


---
