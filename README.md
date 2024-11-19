=======
# Project Xcommerce

One Paragraph of project description goes here

# Project Phase
- Phase 1: Basic monolith for core functions
### Core Functions to Implement (Phase 1: Basic Monolith)

1. **User Authentication & Authorization**
   - User Registration: Implement the ability for users to create an account with email, password, and profile information.
   - Login: Validate credentials and issue JWT tokens for authentication.
   - Authorization: Implement role-based access control (RBAC) for users (e.g., admin, customer).
   - Password Management: Securely store and compare passwords (e.g., bcrypt hashing).

2. **Product Management**
   - CRUD Operations: Create, Read, Update, and Delete products with details such as name, description, price, stock, and category.
   - Product Search: Implement search functionality based on filters like price range, name, or category.

3. **Category Management**
   - CRUD Operations: Manage categories for products (e.g., electronics, clothing).
   - Category Hierarchy: Support for nested categories or subcategories.

4. **Shopping Cart**
   - Add/Remove Items: Allow users to add or remove items from their cart.
   - Cart Persistence: Store cart items between sessions using either a session or persistent storage.
   - Calculate Total: Compute the total price, applying any discounts or promotions.

5. **Order Management**
   - Order Creation: Allow users to place orders based on their cart items.
   - Order Status: Implement different order statuses (e.g., pending, shipped, delivered).
   - Order History: Allow users to view their past orders.

6. **Payment Integration (Optional for Phase 1)**
   - Payment Gateway Integration: Integrate with a payment gateway (e.g., Stripe, PayPal) for processing payments.

7. **Email Notification**
   - Registration Email: Send a confirmation email to users after successful registration.
   - Order Confirmation Email: Send order details and confirmation to users.

8. **Rate Limiting**
   - Implement rate limiting to prevent abuse of your API (e.g., prevent brute force attacks during login).

9. **Testing**
   - Unit Tests: Write tests for core logic (e.g., user authentication, cart calculations).
   - Integration Tests: Test integration with database and external services (e.g., email, payment gateway).
   - Mocking Services: Use mocking frameworks to simulate external services during tests.

10. **Documentation**
    - Swagger API Documentation: Provide interactive API documentation for all endpoints.
    - Project Documentation: Write high-level documentation to explain project architecture, models, and functionality.

- Phase 2: Refactor to Clean Architecture:
### 1. User Registration & Authentication (JWT + OAuth Integration)
- **Description**: Users should be able to register with email/password, verify email, and login using JWT for session management. Implement OAuth for third-party login (e.g., Google, Facebook).
- **Challenge**: Implement strong password policies, user verification with email, and support for both JWT and OAuth.

### 2. Product Inventory Management (CRUD + Stock Levels)
- **Description**: Implement a product management system where admins can create, update, and delete products. Implement a stock management system for products, which should decrease when items are ordered.
- **Challenge**: Handle product data, manage product categories, images, and stock levels. Ensure stock is updated correctly after each order.

### 3. Cart Management (Add, Update, Remove Items)
- **Description**: Implement a cart system where users can add, update quantities, and remove items from their cart. Each cart is linked to a user (or anonymous session).
- **Challenge**: Ensure that cart changes are stored correctly and that users can update cart contents even if they are not logged in.

### 4. Order Processing (Checkout, Payment Gateway Integration)
- **Description**: Implement the order processing flow, where users can complete their checkout process, select a payment method, and submit the order. Integrate with payment gateways (e.g., Stripe, PayPal).
- **Challenge**: Ensure that payments are securely processed and orders are stored with all relevant information (user, products, shipping details).

### 5. Shipping and Address Management
- **Description**: Users should be able to add multiple shipping addresses to their profiles. During checkout, users can select a shipping address and see estimated shipping costs.
- **Challenge**: Manage multiple addresses per user, compute shipping costs, and validate address information.

### 6. Order History & User Profile
- **Description**: Users should be able to view their order history, update their profile, and track the status of their orders.
- **Challenge**: Ensure users can query their order history, track current orders, and edit profile details.

### 7. Admin Dashboard for Order & Product Management
- **Description**: Admin users should have a dashboard to view all orders, manage products, and see inventory levels. Implement basic admin authentication and authorization.
- **Challenge**: Create a separate access layer for admin users to manage orders and products.

### 8. Discounts, Coupons, and Promotions
- **Description**: Implement a discount system where users can apply discount coupons to their orders. Implement promotional pricing for certain products (e.g., seasonal sales).
- **Challenge**: Ensure discounts and promotions are applied correctly and calculate new total prices after discount application.

### 9. Inventory Notification for Low-Stock Items
- **Description**: Implement automatic notifications when products are low in stock (e.g., send email to admin when stock falls below a threshold).
- **Challenge**: Create background jobs that trigger notifications for low-stock items and manage the workflow.

### 10. Reviews and Ratings for Products
- **Description**: Allow users to leave reviews and ratings on products they have purchased.
- **Challenge**: Implement product review functionality with validation for duplicate or invalid reviews.

- Phase 3: Refactor to microservices

## Phase 3: Scalability and Microservices

### 1. Microservices Setup
- **Description**: Split your monolithic application into distinct services (e.g., user service, order service, product service). Ensure they communicate through APIs or a messaging system like Kafka or RabbitMQ.

### 2. API Gateway
- **Description**: Use an API Gateway to route requests to different services and handle cross-cutting concerns like rate limiting, authentication, and authorization.

### 3. Service Discovery
- **Description**: Implement service discovery with tools like Consul or Kubernetes to automatically detect and register services for dynamic scaling.

### 4. Container Orchestration
- **Description**: Use Docker Compose, Kubernetes, or Docker Swarm to manage multiple services, ensuring they can scale horizontally and recover from failures.

### 5. CI/CD Pipelines
- **Description**: Set up continuous integration and deployment (CI/CD) pipelines for automated testing, building, and deploying your microservices.

### 6. Database Sharding
- **Description**: Implement database sharding or partitioning to distribute the data across multiple databases, improving performance and scalability.

### 7. Message Queues
- **Description**: Introduce message queues like RabbitMQ or Kafka to decouple services and enable asynchronous processing of tasks.

### 8. Distributed Tracing
- **Description**: Implement distributed tracing (e.g., using OpenTelemetry) to monitor and trace requests across microservices.

### 9. API Versioning
- **Description**: Handle backward compatibility by implementing versioning strategies for your APIs, ensuring smooth transitions as you scale.

### 10. Load Balancing and Auto-Scaling
- **Description**: Configure load balancers to distribute traffic across multiple instances of each service and implement auto-scaling based on traffic demand.



## Core Model
- User
- Product
- Category
- Order
- Cart


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```