# Programming Workshop Assignments

This document outlines a series of programming workshops designed to build a complete frontend application, from initial design to final refactoring.

## Workshop Overview

| Workshop | Topic | Description |
|----------|-------|-------------|
| Workshop 1 | CSS & JavaScript Animation | Create custom CSS and JavaScript animations for a given HTML file. |
| Workshop 2 | AI-Powered Prototyping | Build and deploy a UI prototype from a Product Requirements Document (PRD) using an AI tool. |
| Workshop 3 | Frontend Development | -- need text -- |
| Workshop 4 | Backend Development & Testing | Design and implement backend services with proper database integration, containerization, and automated testing workflows. |
| Workshop 5 | Code Refactoring & Guidelines | Collaboratively define an AI system prompt to establish a coding style guide and refactor the codebase. |

## Workshop 1: CSS & JavaScript Animation

**Objective**: Animate a static HTML page for KBTG using custom CSS and JavaScript.

**Tasks**:
- You will be given a pre-existing `index.html` file.
- Your task is to write CSS and JavaScript to add animations at the specified points within the file.

## Workshop 2: AI-Powered Prototyping

**Objective**: Create and deploy a UI Proof-of-Concept (POC) from a Product Requirements Document (PRD).

**Tasks**:
- Build the user interface on an AI platform (e.g., v0.dev, lovable.dev, bolt.new) based on the provided PRD.
- Deploy the generated UI to a public hosting service.

## Workshop 3: Frontend Development

**Objective**: Build a complete frontend application using React and modern frontend tooling. Emphasize the use of design tokens to ensure visual consistency across pages and prepare the project for scalability.

**Tasks**:
- Setup a React project with TailwindCSS and React Router through AI-assisted scaffolding.  
- Define design tokens through a UI design system.  
- Create new pages using the design tokens to maintain a unified style.  

### Setup Guide

-   **Vite > React (Typescript)**: https://vite.dev/guide/
-   **Tailwind + Vite**: https://tailwindcss.com/docs/installation/using-vite
-   **Storybook**: https://storybook.js.org/docs/get-started/frameworks/react-vite
-   **Interaction Test**: https://storybook.js.org/docs/writing-tests/interaction-testing
-   **React Router**: https://reactrouter.com/start/declarative/installation

## Workshop 4: Backend Development & Testing

**Objective**: Build and test a backend system that supports the frontend application. Focus on API design, database integration, containerization, and ensuring code quality through unit, integration, and end-to-end testing.

**Tasks**:
- Initialize the backend project with the help of AI (Go + Fiber, Python + FastAPI, or any preferred stack).  
- Convert the UI specification into Swagger documentation, database schema, and backend code (using docs to guide development).  
- Create `Dockerfile` and `docker-compose` for configuration coverage.  
- Implement unit tests with code coverage reports.  
- Perform integration tests to validate database interactions.  
- Conduct end-to-end testing using Playwright, leveraging MCP for analysis support.  

### Setup Guide

- **Go + Fiber** : https://docs.gofiber.io/
- **Playwright**: https://playwright.dev/docs/intro#whats-installed

## Workshop 5: Code Refactoring & Guidelines

**Objective**: To refactor the existing codebase and establish a consistent coding style by collaboratively creating a system prompt for an AI.

**Tasks**:
- As a group, analyze the current project's codebase.
- Collaboratively define a `system prompt` that establishes a clear and effective style guideline for the project.
- The goal is for everyone to contribute ideas to refine this prompt, making the code better and more consistent.

**Repository**:
- Repo for task: https://github.com/mikelopster/kbtg-be-go-lab
- MCP install in vs code: https://code.visualstudio.com/docs/copilot/customization/mcp-servers