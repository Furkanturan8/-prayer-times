# About Project

Hello everyone! The project is a prayer times web application. The purpose here is for users to see daily/monthly prayer times.

## Description
In the prayer times application, vue js is used on the frontend side and golang is used on the backend side. Users can see the daily and monthly prayer times in whichever city they live in. They can also see the remaining time for the next time. With the contact page, users can send feedback or requests to the administrator via e-mail.

## API Endpoints
The system exposes the following API endpoints:

### models/cities
GET /cities - Get all cities <br>

### models/prayer-times
GET /prayer-times/:city - Get prayer times by city for a month<br>
GET /prayer-times/:city/:dayNumber - Get prayer times by city for that(dayNumber) day<br>

### models/phrases
GET /phrases - Get all motto/phrases <br>
GET /phrases/:id - Get a motto/phrase by ID <br>

### models/contact
POST /contact - Create/Send a new contact message<br>



