# Backend Ramadan Dishes Challenge

## Overview

Welcome to the Ramadan Dishes Challenge! This project aims to help Muslims plan their cooking schedule during the holy month of Ramadan based on the five Salawat system. The application will provide information on when to start cooking a specific dish and suggest a dish based on given ingredients and days.

## Features

1. **Cooking Schedule Recommendation:**
   - Determine whether to start cooking a dish before or after *Asr* based on the provided input.
   - The application should consider the cooking time for each dish to make the recommendation.

2. **Dish Suggestion:**
   - Provide a feature to suggest a dish based on given days.

## Technologies Used

- **Language:** Go (Golang)
- **Dependencies:**
    - [gorilla/mux](https://github.com/gorilla/mux) 

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Hamed-amraoui/ramadan-dishes-challenge.git
   cd ramadan-dishes-challenge
2. **Install Dependencies:**
   ```bash
   go get github.com/gorilla/mux
3. **Run the Application:**   
   ```bash
   go build
   .\ramadan.exe

## Usage   

1. **Endpoint 1: Cooking Time**
   - The endpoint `/cooktime` reads two query params, which are:
    - **ingredient**: The required ingredient in the dish.
    - **day**: The *Ramadan* day in which the dish will be cooked
   - The endpoint responds with:
    - The dishes that contain the given ingredient.
    - With each dish, the relative time to *Asr* prayer at which we should start cooking to have it ready 15 minutes before Maghrib prayer. The response should be in the exact following format: `X minutes before/after Asr`
   **Example:**
    ```json
    // GET http://localhost:3000/cooktime?ingredient=Tuna&day=13
    // Response:
   [
      {
         "name": "Brik",
         "ingredients": [
            "Malsouqa",
            "Egg",
            "Tuna"
         ],
         "startcookingtime": "121 minutes After Asr"
      },
      {
         "name": "Blunkett Salad",
         "ingredients": [
            "Bread",
            "Horseradish",
            "Egg",
            "Tuna"
         ],
         "startcookingtime": "116 minutes After Asr"
      }
   ]
2. **Endpoint 2: Suggestions**
   - The endpoint `/suggest` reads one query param:
    - **day**: The *Ramadan* day in which the dish will be cooked
   - The endpoint responds with one dish with the same format as the previous endpoint.
      **Example:**
    ```json
    // GET http://localhost:3000/suggest?day=13
    // Response:
   {
   "name": "Mermez",
   "ingredients": [
      "Chickpea",
      "Tomatoe Paste",
      "Meat",
      "Onion"
   ],
   "startcookingtime": "36 minutes After Asr"
   }
