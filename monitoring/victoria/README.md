## Victoria Metrics Simple Bench

### What is it for?  
This setup provides a simple testing environment to create and catch alerts. It’s particularly useful for developers who are:  
- Writing applications that expose metrics for scraping.  
- Testing alert hooks (e.g., Telegram, email, or custom integrations).  

### Components  
The setup includes a standard `docker-compose.yaml` file with the following services:  
- **VictoriaMetrics**: The database for storing metrics.  
- **vmalert**: For evaluating rules and firing alerts.  
- **Alertmanager**: For managing alert notifications.  
- **Metrics Tester**: A simple app for generating metrics.  
- **Metrics Catcher**: A simple app for catching and displaying alerts.  

With these components, you can easily spin up a complete alert testing environment. Don’t forget to include the necessary configuration files:  
- `prometheus.yml`: Configures metrics scraping.  
- `alerts.rules.yml`: Contains alerting rules for `vmalert`.  
- `alertmanager.yml`: Configures alert routing and notifications.

### How to Use It  
To use this setup, ensure you have Docker Compose installed.  

**Start the environment:**  
```bash
docker compose up -d
```

**Stop the environment:**  
```bash
docker compose down
```

To test alerts, you can use:  
```bash
curl 127.0.0.1:9090/test
```

This will increase the `test_metric_total` counter and trigger an alert. That’s it—your alert system is ready for testing!

### What Else Can Be Done  
You can extend this setup as needed. For example:  
- Customize the `metrics-tester` or `metrics-catcher` applications.  
- Add new alert rules or experiment with different notification channels.  
- Simulate more complex scenarios with multiple metrics and alerts.  