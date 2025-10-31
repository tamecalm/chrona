# Chrona 🌍

[![Go Version](https://img.shields.io/github/go-mod/go-version/golang/go?filename=go.mod&style=flat-square)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

A modern, interactive command-line application built with Go that provides a real-time timezone clock with integrated weather information. Chrona automatically detects your location or allows you to manually select from a comprehensive list of cities worldwide, delivering accurate time and current weather conditions in a beautifully formatted terminal interface.

![App Screenshot](images/demo.png)

## ✨ Features

-   **Real-Time Clock:** A continuously updating digital clock for any selected timezone.
-   **Automatic Geolocation:** Instantly detects your country and region using your IP address for quick timezone lookups.
-   **Manual Selection:** An interactive, user-friendly menu to browse and select any country and major city.
-   **Integrated Weather:** Displays current weather conditions, temperature (in °C and °F), sunrise/sunset times, and whether it's day or night.
-   **Cross-Platform:** Built with Go, it runs seamlessly on Windows, macOS, and Linux.
-   **Elegant UI:** Utilizes colored and styled text to present information in a clear and aesthetically pleasing way.

## 🛠️ Technologies Used

| Technology | Purpose |
| :--- | :--- |
| **[Go](https://golang.org/)** | Core language for building the concurrent and performant CLI application. |
| **[survey/v2](https://github.com/AlecAivazis/survey/v2)** | Powers the interactive prompts and selection menus. |
| **[fatih/color](https://github.com/fatih/color)** | Used for styling the terminal output with colors and attributes. |
| **[ipinfo/go](https://github.com/ipinfo/go)** | Enables automatic location detection based on IP address. |
| **[Open-Meteo API](https://open-meteo.com/)** | Provides real-time weather data for the selected location. |
| **[WorldTimeAPI](https://worldtimeapi.org/)** | Used for fetching accurate timezone information. |

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

-   Go version 1.21 or later installed. You can download it [here](https://golang.org/dl/).

### Installation

1.  **Clone the Repository**
    Open your terminal and clone this repository:
    ```bash
    git clone https://github.com/tamecalm/chrona.git
    ```

2.  **Navigate to the Project Directory**
    ```bash
    cd chrona
    ```

3.  **Install Dependencies**
    Go will handle dependencies automatically, but you can ensure they are up-to-date:
    ```bash
    go mod tidy
    ```

4.  **Build the Application**
    Compile the source code to create an executable file:
    ```bash
    go build -o chrona ./cmd/chrona.go
    ```

## 🖥️ Usage

Once the application is built, you can run it directly from your terminal.

1.  **Execute the Binary**
    ```bash
    ./chrona
    ```
2.  **Location Detection**
    -   You will first be asked if you want to auto-detect your location.
    -   If you choose `Yes`, the application will attempt to find your local timezone.
    -   If you choose `No` or if auto-detection fails, you will be presented with a manual selection menu.

3.  **Manual Selection**
    -   Select a country from the list.
    -   If the country has multiple timezones, select a specific state or major city.

4.  **Real-Time Display**
    -   The clock and weather information for your selected location will be displayed and will update every second.
    -   Press `B` at any time to return to the main country selection menu.
    -   Press `Q` to quit the application.

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  **Fork the Project**
2.  **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4.  **Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **Open a Pull Request**

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

Connect with me on other platforms!

-   **Twitter**: `[@tamecalm]`
