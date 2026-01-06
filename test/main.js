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
async function createProduct() {
  const product = {
    pname: "oooommmmm yes",
    pdescription: "A luxurious modern villa with 4 bedrooms",
    plandetial: "Open floor plan with large windows",
    whatisincluded: ["Foundation", "Walls", "Roof"],
    whatisnotincluded: ["Furniture", "Landscaping"],
    price: 250000,
    dem: {
      wi: 20,
      hi: 10,
      dep: 15
    },
    archistyles: [
      { aname: "Modern" },
      { aname: "Contemporary" }
    ],
    ceilings: [
      { cname: "Living Room Ceiling", ctype: "Vaulted", hi: 12 }
    ],
    garages: [
      { ty: "Attached", entrylocation: "Side", garea: 400, car: 2 }
    ],
    roofdetails: [
      { detail: "Material", val: "Metal Roof" }
    ],
    specialfeatures: [
      { sname: "Smart Home" },
      { sname: "Solar Panels" }
    ]
  };

  try {
    const res = await fetch(`${BASE_URL}/products`, {
      method: "POST",
      headers: { 
        "Content-Type": "application/json",
        "Authorization": `Bearer ${authToken}`
      },
      body: JSON.stringify(product),
    });
    const data = await res.json();
    console.log("Product created:", data);
    return data.pid; // Return product ID for other tests
  } catch (err) {
    console.error("Create product error:", err);
  }
}


// Run the test sequence
(async () => {
  // Optional: signup first if user does not exist
  // await testSignup();


  await testSignup();
  await testLogin();
  await testMe(); 
  await createProduct()
})();
