# Tabi - Technical Specification

## Contents
1. [Introduction](#introduction)
2. [System Architecture](#system-architecture)
3. [Data Models & Database Schema](#data-models--database-schema)
   - [User Model](#user-model)
   - [Trip Model](#trip-model)
   - [Additional Models](#additional-models)
4. [API Endpoints](#api-endpoints)
5. [Business Logic & Rules](#business-logic--rules)
6. [Security Considerations](#security-considerations)
7. [Deployment & CI/CD](#deployment--cicd)
8. [Testing Strategy](#testing-strategy)
9. [Future Enhancements](#future-enhancements)
10. [Appendix & References](#appendix--references)

---

## 1. Introduction
- **Project Name:** Tabi - Trip Planner
- **Version:** 1.0  
- **Author:** Mark Wilson
- **Date:** 08/03/2025
- **Purpose:** Detailed technical design and implementation plan for Tabi's backend, including API endpoints, data models, and integration with PostgreSQL.

---

## 2. System Architecture
**Tech Stack:**  
  - Frontend: Flutter (Dart)  
  - Backend: Go (Gin framework)  
  - Database: PostgreSQL / gORM
  - Authentication: OAuth2  
  - Deployment: Self-hosted in Docker  
  - Logging & Monitoring: Sentry / Zap
**API Stack:** 
  - Amadeus for Flight/Hotel search (https://developers.amadeus.com/)
  - Google places for recommendations (https://developers.google.com/maps/documentation/places/web-service/overview)
  
---

## 3. Data Models & Database Schema

### User Model
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    auth_provider VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Trips Model
```sql
CREATE TABLE trips (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id),
	title VARCHAR(255) NOT NULL,
	destination_country VARCHAR(255),
	destination_city VARCHAR(255) NOT NULL,
	notes TEXT,
	date_from DATETIME NOT NULL,
	date_to DATETIME NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Trip_Flights Model
```sql
CREATE TABLE trip_flights (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	trip_id UUID NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
	type VARCHAR(50) CHECK (type IN ('inbound', 'outbound')) NOT NULL, -- inbound or outbound
	flight_no VARCHAR(50) NOT NULL,
	airline VARCHAR(255),
	cost NUMERIC(10,2),
	passenger_count INT CHECK (passenger_count > 0), -- >= 1
	departure_date TIMESTAMP NOT NULL,
	arrival_date TIMESTAMP NOT NULL CHECK (arrival_date > departure_date),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	notes TEXT
);
```
### Trip_Hotels Model
```sql
CREATE TABLE trip_hotels (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	trip_id UUID NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
	type VARCHAR(255), -- accomodation type, hotel, bnb etc.
	name VARCHAR(500) NOT NULL,
	booking_ref VARCHAR(100),
	booking_agency VARCHAR(100),
	booking_url VARCHAR(500),
	address_country VARCHAR(255) NOT NULL,
	address_city VARCHAR(255) NOT NULL,
	address_line1 VARCHAR(500) NOT NULL,
	address_line2 VARCHAR(500),
	address_line3 VARCHAR(500),
	postcode VARCHAR(50) NOT NULL,
	contact_no VARCHAR(100) NOT NULL,
	contact_name VARCHAR(500),
	contact_email VARCHAR(500),
	date_from TIMESTAMP NOT NULL,
	date_to TIMESTAMP NOT NULL,
	room_type VARCHAR(100),
	check_in_time TIME,
	check_out_time TIME,
	cost_per_night NUMERIC(10,2),
	total_cost NUMERIC(10,2),
	guest_count INT CHECK (guest_count > 0), -- >= 1	
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	notes TEXT	
);
```
