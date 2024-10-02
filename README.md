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
- [Commands](#commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)


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

-go run github.com/tmc/langchaingo/examples/ollama-completion-example@main

You should receive an output similar to this:

- The first human to set foot on the moon was Neil Armstrong, an American astronaut, who stepped onto the lunar surface.

those are the instactions from :

[LangChainGo](https://tmc.github.io/langchaingo/docs/getting-started/guide-ollama)
---
