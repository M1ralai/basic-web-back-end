function get1() {
  fetch("http://localhost:8080/api/users/32", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: "Atahan Demirer",
      password: "Atahan Demirer",
    }),
  })
    .then((response) => {
      console.log(response)
      if (!response.ok) {
        throw new Error("Request başarısız: " + response.status);
      }
      return response.json();
    })
    .then((data) => console.log("Gelen veri:", data))
    .catch((error) => console.error("Hata:", error));
}


function get2() {
    fetch("http://localhost:8080/api/users", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      password: "Atahan Demirer",
    }),
  })
    .then((response) => {
      console.log(response)
      if (!response.ok) {
        throw new Error("Request başarısız: " + response.status);
      }
      return response.json();
    })
    .then((data) => console.log("Gelen veri:", data))
    .catch((error) => console.error("Hata:", error));
}

function get3() {
    fetch("http://localhost:8080/api/users", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
        username: "Atahan Demirer",
        password: "Atahan Demirer",
    }),
  })
    .then((response) => {
      console.log(response)
      if (!response.ok) {
        throw new Error("Request başarısız: " + response.status);
      }
      return response.json();
    })
    .then((data) => console.log("Gelen veri:", data))
    .catch((error) => console.error("Hata:", error));
}
