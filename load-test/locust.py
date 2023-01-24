from locust import HttpUser, task, between

class QuickstartUser(HttpUser):

    @task
    def sample_call(self):
        self.client.get("/")
        self.client.get("/ping")
