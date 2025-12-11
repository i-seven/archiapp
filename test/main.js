const BASE_URL = "http://localhost:8080";

let authToken = "";

// Signup a new user
async function testSignup() {
  const user = {
    email: "savi.543@gmail.com",
    password: "Uyjhmn65@!",
    name: "iz7zi"
  };

  try {
    const res = await fetch(`${BASE_URL}/users/signup`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });
    const data = await res.json();
    console.log("Signup response:", data);
  } catch (err) {
    console.error("Signup error:", err);
  }
}
async function testLogin() {
  const res = await fetch('http://localhost:8080/users/login', {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email: "savi.543@gmail.com",
      password: "Uyjhmn65@!"
    })
  });
  const data = await res.json();
  console.log("Login response:", data);
  authToken = data.token;
}

// Get current logged-in user
async function testMe() {
  try {
    const res = await fetch(`${BASE_URL}/users/me`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${authToken}`,
      },
    });
    const data = await res.json();
    console.log("Current user (/me):", data);
  } catch (err) {
    console.error("Me error:", err);
  }
}

// Run the test sequence
(async () => {
  // Optional: signup first if user does not exist
  // await testSignup();


  await testSignup();   // login and set authToken
  await testLogin();   // login and set authToken
  await testMe();      // call protected route using token
})();
