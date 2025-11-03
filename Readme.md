# Fundamentals of System Design 

---

### What is System Design?

**Answer:**

System Design is the process of defining the **architecture, components, and data flow** of a system to meet both **functional** and **non-functional** requirements (like scalability, reliability, and performance).

In simple terms, it’s about **creating the blueprint of a large-scale system** — how different parts (APIs, databases, services, caches, etc.) interact to serve millions of users efficiently.

**Key Components:**

1. **Functional Requirements** – What the system should do.
2. **Non-Functional Requirements** – How well it should perform (scalability, latency, fault tolerance).
3. **High-Level Design** – System architecture, data flow, service interactions.
4. **Low-Level Design** – Data models, algorithms, and detailed interactions between components.


---

### Why is System Design Important?

**Answer:**

System Design is important because it ensures that a system can **scale, perform, and remain reliable** as it grows. It helps engineers make the right **architectural decisions early**, preventing costly rework later.

**Key Reasons:**

1. **Scalability** – Handles increased user load efficiently.
2. **Reliability** – Ensures the system stays available even if parts fail.
3. **Performance** – Maintains low latency and fast response times.
4. **Maintainability** – Makes the system easier to update and debug.
5. **Cost Efficiency** – Optimizes infrastructure and operational costs.

---
## Network & Communication

### Why Does Networking Matter in System Design?

**Answer:**

Networking is the **backbone of distributed systems** — every service-to-service call, API request, and database query relies on network communication. Understanding networking helps you design systems that are **fast, reliable, and fault-tolerant**.

**Key Reasons:**

1. **Service Communication** – Enables interaction between microservices, databases, and clients.
2. **Performance Optimization** – Helps reduce latency and bandwidth usage.
3. **Scalability** – Supports horizontal scaling using load balancers and distributed nodes.
4. **Fault Tolerance** – Proper network design avoids single points of failure.
5. **Security** – Ensures safe data transfer using encryption and secure protocols (HTTPS, TLS).

---

### How Does Networking Impact Large-Scale Systems?

**Answer:**

In large-scale systems, **networking directly affects performance, scalability, and reliability**. As the system grows, data must travel across multiple servers, regions, and services — making network efficiency critical.

**Key Impacts:**

1. **Latency** – Each network call adds delay; optimizing connections reduces response time.
2. **Bandwidth Usage** – Inefficient data transfer can slow down or overload the system.
3. **Scalability** – Proper network design (load balancers, CDNs, replication) enables horizontal scaling.
4. **Fault Tolerance** – Redundant paths and retries help handle node or region failures.
5. **Consistency** – Network delays can cause data synchronization issues across distributed nodes.
6. **Cost** – Cross-region data transfers and high bandwidth usage increase infrastructure costs.

---

## 🌐 Understanding of IP Address

---

### Introduction to IP Address

**Answer:**

An **IP (Internet Protocol) Address** is a unique identifier assigned to every device connected to a network. It allows devices (like servers, clients, routers) to **find and communicate** with each other over the internet or a local network.

**Key Points:**

1. **Purpose** – Identifies and locates devices in a network.
2. **Format** – Can be **IPv4** (e.g., `192.168.1.1`) or **IPv6** (e.g., `2001:0db8::1`).
3. **Types** –

   * **Public IP:** Used for communication over the internet.
   * **Private IP:** Used within internal/local networks.
4. **Static vs Dynamic:**

   * **Static IP** – Manually assigned, doesn’t change.
   * **Dynamic IP** – Assigned automatically by DHCP, can change over time.

---

### What is IPv4?

**Answer:**

**IPv4 (Internet Protocol version 4)** is the **fourth version** of the Internet Protocol and the most widely used system for identifying devices on a network. It uses a **32-bit address format**, allowing around **4.3 billion unique addresses**.

**Key Points:**

1. **Format:** Four numbers separated by dots (e.g., `192.168.1.1`).
2. **Address Range:** 0.0.0.0 to 255.255.255.255.
3. **Address Space:** ~4.3 billion addresses (2³²).
4. **Representation:** Each part (octet) ranges from 0–255.
5. **Limitation:** Due to internet growth, IPv4 addresses are running out — leading to **IPv6** adoption.

---

### What is IPv6?

**Answer:**

**IPv6 (Internet Protocol version 6)** is the **successor to IPv4**, developed to overcome the limitation of IPv4’s address exhaustion. It uses a **128-bit address format**, providing an almost unlimited number of unique IP addresses.

**Key Points:**

1. **Format:** Eight groups of four hexadecimal digits separated by colons (e.g., `2001:0db8:85a3:0000:0000:8a2e:0370:7334`).
2. **Address Space:** 2¹²⁸ possible addresses — enough for every device globally.
3. **Simplified Configuration:** Supports auto-configuration without DHCP.
4. **Security:** Built-in support for IPsec (encryption and authentication).
5. **Efficiency:** Improved routing and reduced network congestion compared to IPv4.

---

### Private vs Public IP Address

**Answer:**

**Private IP addresses** are used **within local networks** (like home, office, or internal company systems), while **Public IP addresses** are used to **communicate over the internet**.

| **Aspect**          | **Private IP**                                                                                  | **Public IP**                                 |
| ------------------- | ----------------------------------------------------------------------------------------------- | --------------------------------------------- |
| **Scope**           | Used within a local/private network                                                             | Used across the internet                      |
| **Access**          | Not directly accessible from the internet                                                       | Accessible globally                           |
| **Assigned By**     | Local network devices (e.g., router/DHCP)                                                       | Internet Service Provider (ISP)               |
| **Examples (IPv4)** | `10.0.0.0 – 10.255.255.255`<br>`172.16.0.0 – 172.31.255.255`<br>`192.168.0.0 – 192.168.255.255` | Any address not in the private IP ranges      |
| **Security**        | More secure, isolated from external access                                                      | Exposed to external traffic, needs protection |
| **Use Case**        | Home LANs, corporate intranets, internal servers                                                | Web servers, public APIs, cloud services      |

**In short:**
Private IP = for internal communication.
Public IP = for external (internet) communication.

---

### Why Do We Need Private IPs?

**Answer:**

Private IPs are essential for enabling **internal communication** within local networks **without consuming public IP addresses**. They help build secure, isolated environments for devices to connect and share resources efficiently.

**Key Reasons:**

1. **Address Conservation:** Reduces the need for public IPs, which are limited in IPv4.
2. **Internal Communication:** Allows devices (like laptops, printers, and servers) to connect within the same network.
3. **Security:** Keeps internal systems hidden from the public internet, reducing attack risks.
4. **Cost Efficiency:** No need for every device to have its own public IP.
5. **Network Flexibility:** Enables local routing, DHCP usage, and NAT (Network Address Translation).

**In short:**
Private IPs make internal networking **secure, scalable, and cost-effective** without exposing internal systems to the public web.

---

### The Role of IP in System Design

**Answer:**

IP (Internet Protocol) plays a fundamental role in system design by enabling **communication between different components** of a distributed system — servers, databases, load balancers, and clients — across networks.

**Key Roles:**

1. **Identification:** Each device or service in a network is uniquely identified by an IP address.
2. **Routing:** IP ensures that data packets reach the correct destination across local or global networks.
3. **Scalability:** Allows horizontal scaling by assigning unique IPs to new servers or microservices.
4. **Load Balancing:** Load balancers use IPs to route traffic to multiple backend servers.
5. **Security & Isolation:** Private IPs separate internal traffic from public exposure, improving security.
6. **High Availability:** Enables failover and redundancy through multiple IP-based routes and replicas.

**In short:**
IP addresses are the **foundation of connectivity** in any large-scale system — making communication, scalability, and reliability possible.

---

## 🌐 How DNS Works

---

### Introduction to DNS

**Answer:**

**DNS (Domain Name System)** is the **internet’s phonebook**, translating **human-readable domain names** (like `google.com`) into **IP addresses** (like `142.250.190.78`) that computers use to identify each other on a network.

Without DNS, users would have to remember complex IP addresses to access websites.

**Key Points:**

1. **Purpose:** Converts domain names to IP addresses for easy access.
2. **Analogy:** Just like saving a contact name instead of remembering a phone number.
3. **Function:** When you enter a website URL, DNS finds the correct IP so your browser can connect to the right server.
4. **Hierarchy:** DNS operates through a distributed hierarchy of servers — Root, TLD, and Authoritative servers.
5. **Performance:** DNS caching helps speed up future lookups and reduce latency.

**In short:**
DNS makes the internet **human-friendly** by bridging the gap between names and numerical IP addresses.

---

### Types of DNS Servers

**Answer:**

DNS operates through a **hierarchy of servers**, each playing a specific role in resolving a domain name to its IP address. The resolution process usually involves multiple servers working together.

**Main Types of DNS Servers:**

1. **Recursive Resolver:**

   * The first server contacted by your browser or device.
   * Responsible for querying other DNS servers to find the correct IP address.
   * Often managed by ISPs or public DNS providers (e.g., Google DNS `8.8.8.8`).

2. **Root DNS Server:**

   * The starting point of DNS resolution.
   * Directs queries to the correct **Top-Level Domain (TLD)** server (like `.com`, `.org`, `.net`).

3. **TLD (Top-Level Domain) Server:**

   * Manages domain extensions (e.g., `.com`, `.in`, `.org`).
   * Points to the **Authoritative DNS Server** of the specific domain.

4. **Authoritative DNS Server:**

   * The final source of truth for a domain.
   * Holds actual DNS records (like A, AAAA, MX, and CNAME) and provides the corresponding IP address.

**In short:**
The DNS resolution chain works as:
**Client → Recursive Resolver → Root Server → TLD Server → Authoritative Server → IP Address Returned**

---

### DNS Caching and Performance Optimization

---

### Why Caching Matters

**Answer:**

DNS caching is crucial because it **reduces lookup time** and **improves performance** by storing previously resolved domain-to-IP mappings.
Instead of querying DNS servers repeatedly, the system can reuse cached results, leading to **faster response times** and **lower network load**.

**Benefits:**

* Reduces DNS query latency.
* Minimizes load on upstream DNS servers.
* Improves user experience with quicker website access.
* Decreases bandwidth and infrastructure costs.

---

### Where Caching Occurs

**Answer:**

Caching can happen at multiple layers in the DNS resolution chain:

1. **Browser Cache:**
   Stores recent DNS lookups for a short duration to speed up repeated visits.

2. **Operating System (OS) Cache:**
   The OS maintains a local DNS cache shared across applications.

3. **Recursive Resolver Cache:**
   ISPs or public resolvers (like Google DNS, Cloudflare DNS) cache responses to serve multiple users efficiently.

4. **CDN or Proxy Cache:**
   Content Delivery Networks often cache DNS responses close to users to reduce lookup distance.

---

### What is TTL (Time To Live)

**Answer:**

**TTL (Time To Live)** defines how long a DNS record can be stored in a cache before it must be refreshed from the authoritative server.
It’s measured in **seconds** and balances freshness with performance.

**Example:**
If a DNS record has a TTL of `3600`, it means the cached record is valid for **1 hour**.

**Key Insight:**

* **Higher TTL:** Better performance, slower updates.
* **Lower TTL:** Faster updates, more DNS lookups.

---

### The Domain Name Resolution Process — Step-by-step

**Answer:**

1. **User enters URL / app requests domain**
   The client (browser or app) starts by asking the OS to resolve the domain name.

2. **Browser cache check**
   Browser checks its internal DNS cache. If found and not expired, use it and finish.

3. **OS / Local resolver cache**
   If browser cache misses, the OS DNS cache (or local DNS resolver) is checked next.

4. **Hosts file check**
   The OS checks the local `hosts` file for a static mapping. If present, use it.

5. **Query to Recursive Resolver**
   If still unresolved, the OS sends a DNS query to the configured recursive resolver (usually ISP or public DNS like 8.8.8.8). The recursive resolver will perform the full resolution on behalf of the client.

6. **Recursive resolver cache**
   The resolver checks its cache. If a cached answer exists and is valid (TTL not expired), it returns the IP to the client.

7. **Root server query (if no cache)**
   If the resolver has no cached answer, it queries a Root DNS server. The Root responds with the authoritative TLD server location for the domain’s extension (e.g., `.com`).

8. **TLD server query**
   The resolver queries the TLD server returned by the Root. The TLD responds with the authoritative name server(s) for the specific domain.

9. **Authoritative server query**
   The resolver queries the domain’s authoritative DNS server. The authoritative server returns the final DNS record (A, AAAA, CNAME, etc.) with its TTL.

10. **Resolver returns result to client**
    The recursive resolver caches the answer (for TTL seconds) and returns the IP address to the client.

11. **Client caches the result**
    The OS and browser cache the record according to TTL for future requests.

12. **Client connects to server**
    With the IP known, the client initiates a TCP/TLS connection (e.g., TCP handshake, TLS handshake) to the server and requests the resource (HTTP request).

13. **Subsequent requests**
    Subsequent lookups follow the same flow but often end at cache layers (browser, OS, resolver) until TTL expiry.

Notes:

* Caching at multiple layers reduces lookup time and network load.
* Recursive resolvers can perform iterative queries: they query Root → TLD → Authoritative and assemble the answer.
* Modern variants include DNS over HTTPS (DoH) and DNS over TLS (DoT), which encrypt DNS queries between client and resolver.

---

### Importance of DNS in Large-Scale Systems

**Answer:**

DNS plays a **critical role in large-scale systems** by ensuring smooth routing, scalability, and high availability across distributed infrastructure. It acts as the **entry point** for almost every user request to reach the right server or service.

**Key Reasons:**

1. **Traffic Routing:**
   DNS directs user requests to appropriate servers, data centers, or regions based on geography, latency, or load.

2. **Scalability:**
   Enables horizontal scaling by balancing traffic across multiple IPs or service instances.

3. **High Availability:**
   Allows failover by rerouting traffic to healthy servers or backup regions when one fails.

4. **Load Balancing Integration:**
   Works with DNS-based load balancing (like round-robin or geo-DNS) to distribute requests efficiently.

5. **Performance Optimization:**
   DNS caching reduces resolution time, improving user experience for global systems.

6. **Service Discovery:**
   Helps microservices find each other dynamically through DNS-based service registries.

7. **Security and Control:**
   Supports protection via DNS filtering, DNSSEC, and traffic management policies.

**In summary:**
DNS is not just a name-to-IP translator—it’s a **strategic component** in system design that ensures **speed, resilience, and global scalability** in modern distributed systems.

---
