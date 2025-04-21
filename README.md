<p align="center">
  <h1>Book Ease</h1>
  <p><em>Your all‑in‑one travel booking platform: flights, hotels, packages.</em></p>
</p>

---

## 📖 Table of Contents

1. [Project Overview](#project-overview)  
2. [Team & Responsibilities](#team--responsibilities)  
3. [Tech Stack](#tech-stack)  
4. [Key Features](#key-features)  
5. [Architecture & API](#architecture--api)  
6. [Getting Started](#getting-started)  
   - [Prerequisites](#prerequisites)  
   - [Installation](#installation)  
   - [Configuration](#configuration)  
   - [Running the App](#running-the-app)  
7. [Testing](#testing)  
8. [Coding Standards](#coding-standards)  
9. [Contributing](#contributing)  
10. [License & Acknowledgments](#license--acknowledgments)  

---

## Project Overview

Book Ease is a single‑page application that streamlines travel planning by unifying flight, hotel, and package bookings. Built for seamless user experience, it provides:

- **Centralized Search** across providers  
- **Real‑time Pricing** with package + accommodation breakdown  
- **Secure Authentication** and session management  
- **Responsive Design** for desktop and mobile  

---

## Team & Responsibilities

| Member              | Role         | Responsibilities                                    |
|---------------------|--------------|-----------------------------------------------------|
| Prathima Dodda      | Frontend Lead| Component design, routing, payment integration      |
| Sai Preethi Kota    | Frontend     | Authentication flows, form validation, styling      |
| Varshini Kopparla   | Backend Lead | Go‑based API endpoints, database schema, migrations |
| Karthik Karnam      | Backend      | Accommodation controllers, unit tests, validation   |

---

## Tech Stack

- **Frontend:** React JS, React Router, Tailwind CSS  
- **Backend:** Go, Gin‑Gonic, MySQL  
- **Testing:** Jest, React Testing Library, Cypress  
- **CI/CD:** GitHub Actions  
- **Lint & Format:** ESLint, Prettier  

---

## Key Features

- **Package Search & Details**  
- **Live Price Calculation:**  
  - `GET /prices?package_id=<id>[&accommodation_id=<id>]`  
- **Hotel Booking Flow** with optional add‑on  
- **User Authentication:** JWT‑based login/register  
- **Global State Management:** React Context & Hooks  
- **Unit & End‑to‑End Tests** coverage ≥ 85%  

---

## Architecture & API

### Frontend

