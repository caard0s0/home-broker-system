<div id="top"></div>

<!-- About the Project -->
<div align="center">
<h2>Home Broker Algorithm - Microservice</h2>
<p>A Home Broker algorithm for Investment Brokers, developed with <a href="https://go.dev/">Go</a>.</p>
<a href="https://github.com/caard0s0/home-broker-system/issues">Report Bugs</a>
&nbsp;&bull;&nbsp;
<a href="https://github.com/caard0s0/home-broker-system/actions">Actions</a>
&nbsp;&bull;&nbsp;
<a href="https://github.com/caard0s0/home-broker-system/pulls">Pull Requests</a>
</div>

&nbsp;

![home_broker_flowchart](https://github.com/caard0s0/home-broker-system/assets/95318788/107c4c7b-e64a-45c0-af85-4fc8a494edf2)

An algorithm intended for share purchase and sale operations in a Home Broker environment. The application was built on a microservices architecture, using Go and Kafka, which is adopted as a messaging system, where purchase and sale orders are produced on a specific topic. The algorithm, in turn, consumes these orders, processes them and, when there is a match, produces the resulting transaction in another topic. all encapsulated in Docker containers with the advantage of easy access to Confluent's Control Center for monitoring and managing Kafka.

&nbsp;

<h3>Built With</h3>

[![Tech Tools](https://skillicons.dev/icons?i=go,docker,kafka)](https://skillicons.dev)


<!-- Table of Contents -->
<details>
<summary>Table of Contents</summary>
<ol>
<li>
    <a href="#getting-started">Getting Started</a>
    <ul>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#usage">Usage</a></li>
        <li><a href="#tests">Tests</a></li>
    </ul>
</li>
<li><a href="#license">License</a></li>
<li><a href="#contact">Contact</a></li>
</ol>
</details>

&nbsp;


<!-- Getting Started -->
<h2 id="getting-started">Getting Started</h2>

<p>To get started, You need to have <strong>Go 1.21+</strong> installed on your machine, for more information visit <a href="https://go.dev/dl/">Go Downloads</a>. You also need to have <strong>Docker Desktop</strong> installed, for more information visit <a href="https://www.docker.com/products/docker-desktop/">Docker Desktop Install</a>.</p>

<p><strong>OBS:</strong> This guide is designed to run this project locally (Local Development), on Linux-based systems.</p>


<!-- Installation -->
<h3 id="installation">Installation</h3>

1. Clone the repository.
```bash
git clone https://github.com/caard0s0/home-broker-system.git
```


<!-- Usage -->
<h2 id="usage">Usage</h2>

<p>After completing the installation, you can run the project.</p>

1. Create and run the <strong>Containers</strong>.

    ```cmd
    docker compose up -d
    ```

2. Open a browser tab at `localhost:9021` to access your <strong>Cluster</strong> in the <strong>Confluent Control Center</strong>.

    ![cluster_control_center](https://github.com/caard0s0/home-broker-system/assets/95318788/9360c92d-06cb-4b80-97f7-9e3f0cbdcc45)

3. Opening your terminal at the root of the application, run the project.

    ```zsh
    go run cmd/main.go
    ```

4. Create a new <strong>Topic</strong>.

    <strong>WARNING:</strong> The <strong>Topic name</strong> must be exactly as written in the image.

    ![new_topic](https://github.com/caard0s0/home-broker-system/assets/95318788/42f84f86-b7b0-4385-8d76-e4f0dcc22427)

5. Create two new messages (<strong>one at a time</strong>) using the data below, to generate a <strong>New Match</strong> and consequently a <strong>Transaction</strong>.

    ```json
    {
        "order_id": 1,
        "investor_id": 1,
        "stock_id": "Stock1",
        "current_shares": 10,
        "shares": 10,
        "price": 10,
        "order_type": "BUY"
    }

    {
        "order_id": 2,
        "investor_id": 2,
        "stock_id": "Stock1",
        "current_shares": 10,
        "shares": 10,
        "price": 10,
        "order_type": "SELL"
    }
    ```

    ![create_message](https://github.com/caard0s0/home-broker-system/assets/95318788/1032bf55-54bc-4d6e-b354-15d0882c7437)

6. Opening your terminal, you will see the result of the transaction exactly as it is below.

    <p><strong>OBS:</strong> You can also see the result in your <strong>Cluster</strong>.</p>

    ```json
    {
        "order_id": 1,
        "investor_id": 1,
        "stock_id": "Stock1",
        "order_type": "BUY",
        "status": "CLOSED",
        "partial": 0,
        "shares": 10,
        "transactions": [
            {
                "transaction_id": "324cb1a1-e73d-4deb-99bc-bf1440907412",
                "buyer_id": 1,
                "seller_id": 2,
                "stock_id": "Stock1",
                "price": 10,
                "shares": 10
            }
        ]
    }
    {
        "order_id": 2,
        "investor_id": 2,
        "stock_id": "Stock1",
        "order_type": "SELL",
        "status": "CLOSED",
        "partial": 0,
        "shares": 10,
        "transactions": [
            {
                "transaction_id": "324cb1a1-e73d-4deb-99bc-bf1440907412",
                "buyer_id": 1,
                "seller_id": 2,
                "stock_id": "Stock1",
                "price": 10,
                "shares": 10
            }
        ]
    }
    ```


<!-- Tests -->
<h2 id="tests">Tests</h2>

<p>To be able to run all the tests, follow the command below.</p>

1. Run all the <strong>Tests</strong>.

```cmd
go test -v -cover ./...
```


<br>


<!-- License -->
<h2 id="license">License</h2>

This project is being distributed under the <strong>MIT License</strong>, see ```LICENSE.txt``` for more information.


<br>


<!-- Contact -->
<h2 id="contact">Contact</h2>

* Software Engineer  
* Vinicius Cardoso - <a href="mailto:cardoso.business.ctt@gmail.com">Email</a>

<p align="right">
<a href="#top"> &uarr; back to top</a>
</p> 
