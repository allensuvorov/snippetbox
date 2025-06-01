# snippetbox

A simple web application built in Go that lets users create, view, and share snippets of text or code.

This project is designed to:
- Showcase idiomatic Go for web applications
- Demonstrate layered architecture with clean separation of concerns
- Serve as a learning base for middleware, session handling, and secure web app development

---

## 📦 Features

- Snippet creation form with expiry options
- View individual snippets via unique URLs
- Homepage with recent snippets list
- User registration and login system
- Secure session management with cookies
- Password hashing with bcrypt
- CSRF protection and input validation
- Graceful error handling and custom logging
- Clean MVC-style project structure

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.13 or later
- MySQL 5.7+ (or compatible)

### 🛠 Setup Instructions

1. **Clone the repository:**

```bash
git clone https://github.com/allensuvorov/snippetbox.git
cd snippetbox
