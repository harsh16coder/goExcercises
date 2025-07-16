# Simple URL Shortener in Go

This project is a **minimal URL shortener** built with Go. The goal is to demonstrate how you can take a long URL from a user, generate a unique short code, and serve a redirect from a local web server—all using Go’s standard library.

---

## 🚀 What Does This Project Do?

- **Reads a URL from the user** via the command line.
- **Generates a unique, random short code** for that URL.
- **Stores the mapping** between the short code and the original URL in memory.
- **Starts a web server on port 8080**.
- **Redirects anyone visiting `http://localhost:8080/<shortcode>`** to the original URL.

---

## 🛠️ How It Works

1. **User Input:**  
   When you run the program, it prompts you to enter a URL.

2. **Short Code Generation:**  
   The program creates a random 6-character code (like `aB3xYz`).

3. **Mapping:**  
   It stores the mapping between the short code and your original URL in memory.

4. **Web Server:**  
   The program starts an HTTP server on `localhost:8080`.  
   If you visit `http://localhost:8080/<shortcode>`, you’ll be redirected to your original URL.

---

## 🏃‍♂️ How to Run

1. **Clone or copy this project to your machine.**

2. **Open a terminal in the project directory.**

3. **Run the program:**
   ```sh
   go run main.go
   ```

4. **Enter a URL** when prompted (e.g., `https://www.google.com`).

5. **Copy the generated short URL** (e.g., `http://localhost:8080/Ab12Cd`) and open it in your browser.  
   You’ll be redirected to the original URL!

---

## 📝 Example

```
Enter the URL: https://www.google.com
Shortened URL: http://localhost:8080/aB3xYz
Server running on http://localhost:8080
```

Now, visiting `http://localhost:8080/aB3xYz` in your browser will take you to `https://www.google.com`.

---

## ⚠️ Notes

- **In-memory only:** All mappings are lost when the program stops.
- **Single URL per run:** This demo only asks for one URL per run, but you can extend it to handle multiple URLs or add a web form.
- **No persistence or authentication:** This is a learning/demo project.

---

## 📚 What You’ll Learn

- How to read user input from the command line in Go.
- How to generate random strings.
- How to use Go’s built-in HTTP server.
- How to handle HTTP requests and perform redirects.

---

## 💡 Next Steps

Want to extend this project? Try:
- Adding a web form for URL submission.
- Storing mappings in a file or database.
- Handling duplicate URLs or short codes.
- Adding analytics for link clicks.

---

Happy coding! 🚀