# About Project

Hello everyone! The project is a prayer times web application. The purpose here is for users to see daily/monthly prayer times.

## Description
In the prayer times application, vue js is used on the frontend side and golang is used on the backend side. Users can see the daily and monthly prayer times in whichever city they live in. The page refreshes every 30 seconds. In this way, they can see the remaining time (counter) for the next time as up-to-date and with a low margin of error. With the contact page, users can send feedback or requests to the administrator via e-mail.

## API Endpoints
The system exposes the following API endpoints:

### models/city
GET /cities - Get all cities <br>

### models/timings
GET /prayer-times/:city - Get prayer times by city for a month<br>
GET /prayer-times/:city/:dayNumber - Get prayer times by city for that(dayNumber) day<br>

### models/phrases
GET /phrases - Get all motto/phrases <br>
GET /phrases/:id - Get a motto/phrase by ID <br>

### models/contact
POST /contact - Create/Send a new contact message<br>

## Video of the project:
https://github.com/user-attachments/assets/1d18e83c-e535-4bb1-b284-5a0437ad7cba


