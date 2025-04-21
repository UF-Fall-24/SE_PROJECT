<p align="center">
  <h1>📖 Book Ease — Travel Booking Web App</h1>
  <p><em>Streamline your journey: search, compare & book flights, hotels & packages in one place.</em></p>
</p>

---

## 🚀 Table of Contents
1. [Project Overview](#project-overview)
2. [Team & Roles](#team--roles)
3. [Tech Stack](#tech-stack)
4. [🎯 Features](#-features)
5. [💻 Getting Started](#-getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
   - [Environment Variables](#environment-variables)
   - [Running Locally](#running-locally)
6. [🧪 Testing](#-testing)
7. [📁 Project Structure](#-project-structure)
8. [🤝 Contributing](#-contributing)
---

## Project Overview
**Book Ease** centralizes travel services—flights, hotels, packages—into a single intuitive platform. Users can:
- 🔍 Search for travel packages and hotels by destination and dates
- 💰 Instantly calculate combined package + accommodation pricing
- 🔒 Securely register & authenticate
- 📱 Enjoy a responsive, mobile‑friendly React UI

This Sprint 4 update includes:
- Complete **homepage redesign** and navigation
- **Dynamic price** API integration (`GET /prices`)
- **Unit tests** for core components & **Cypress E2E** scenarios
- Updated **API documentation** for backend collaborators

---

## Team & Roles
| Member              | Role       | Responsibilities          |
|---------------------|------------|---------------------------|
| Prathima Dodda      | Frontend   | UI components, routing, payment flow integration |
| Sai Preethi Kota    | Frontend   | Authentication, forms, CSS styling |
| Varshini Kopparla   | Backend    | Go API endpoints, DB migrations |
| Karthik Karnam      | Backend    | Accommodation & package controllers, tests |

---

## Tech Stack
| Layer          | Technology                           |
|----------------|--------------------------------------|
| **Frontend**   | React, React Router, Tailwind CSS    |
| **Backend**    | Go (Gin), MySQL                      |
| **Testing**    | Jest, React Testing Library, Cypress |
| **CI/CD**      | GitHub Actions                       |

---

## 🎯 Features
1. **Homepage** with featured packages & search bar
2. **Package Details**: view itinerary, pricing, location
3. **Booking Flow**:
   - **Payment**, passing only `package_id`
   - **Accommodation** selection, then **combined price** lookup
4. **User Auth**: register, login & token‑based sessions
5. **Responsive Design** for mobile & desktop

---

## 💻 Getting Started
### Prerequisites
- Node.js ≥ 18.x & npm ≥ 9.x
- Go ≥ 1.20 & MySQL
- Git

### Installation
```bash
# Clone frontend
git clone https://github.com/UF-Fall-24/SE_PROJECT.git
cd SE_PROJECT/src
npm install

# Clone backend
cd ../SE_PROJECT-BACKEND-BookEase
go mod download
```

### Environment Variables
Create a `.env.local` in `/src/`:
```ini
REACT_APP_API_BASE=http://localhost:8080
```
And a `.env` in backend:
```ini
DB_DSN=username:password@tcp(localhost:3306)/bookease
```

### Running Locally
```bash
# Start backend (port 8080)
go run main.go

# In a new terminal, start frontend (port 3000)
cd SE_PROJECT/src
npm start
```
Visit <http://localhost:3000> and login with sample user.

---

## 🧪 Testing
**Unit Tests (Jest)**
```bash
npm test
```

**E2E Tests (Cypress)**
```bash
npm run cypress:open
# or headless
npm run cypress:run
```

---

## 📁 Project Structure
```
src/
├── assets/            # images & icons
├── components/        # reusable components (Payment, PackageDetails…)
├── hooks/             # custom React hooks
├── pages/             # route-level views
├── services/          # API wrapper modules
├── styles/            # Tailwind overrides & global CSS
├── tests/             # Jest & RTL tests
├── App.jsx            # main router & layout
├── index.js           # React entry point
└── README.md          # this file
```

---

## 🤝 Contributing
1. **Fork** & **clone** the repo
2. Create a feature branch: `git checkout -b feature/new-ui`
3. Commit your changes: `git commit -m "feat: add search filter"`
4. Push: `git push origin feature/new-ui`
5. Open a **Pull Request** against `main`, link relevant issues.

Please keep branches up to date with `main` and ensure all tests pass before requesting review.

---

*Happy coding & safe travels!* 🚀

