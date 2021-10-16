# covid-cases-near-me

An application built in GoLang, that exposes APIs for fetching Covid Cases info from public APIs and persists data in Mongo.
Another API is then used to collect user's location coordinates and uses a reverse geo-coding service to return data.

Source Used: https://covid19api.com/
Assumptions Related to Source: 
-   The data is collated from a data source maintained by John Hopkins University
-   Data is typically updated in an interval of 30 mins. In order to ensure that we donot query large amounts of corrupt data, we always fetch data that was updated 1 hour before the current time. (API Limitations)

Disclaimer:
This is just a prototype working project, there is no guarantee that the API used to fetch covid cases would be accurate.

Major Libraries Used:
Redis -> go-redis
mongoDB -> mongo-go-driver
Framework -> echo
Swagger -> swaggo

Database Constraints:
There exists unique key constraints on the basis of [country,date] for country collection and [province,date] for state collection.
