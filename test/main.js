let authToken = "";

// First, login to get the token
function testLogin() {
  const user = {
    email: "hadimousavi.543@gmail.com",
    password: "Uyjhmn65@!",
  };

  fetch('http://localhost:3000/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(user)
  })
    .then(res => res.json())
    .then(data => {
      console.log('Logged in:', data);
      authToken = data.token; // Save token for later requests
    })
    .catch(err => console.error('Login error:', err));
}

// Create a product (requires auth)
function testCreateProduct() {
  const product = {
    name: "Test Product",
    price: 99.99,
    description: "A product for testing"
  };

  fetch('http://localhost:3000/products', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(product)
  })
    .then(res => res.json())
    .then(data => console.log('Created product:', data))
    .catch(err => console.error('Error:', err));
}

// Update a product (requires auth)
function testUpdateProduct(id) {
  const updated = {
    name: "Updated Product",
    price: 79.99
  };

  fetch(`http://localhost:3000/products/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(updated)
  })
    .then(res => res.json())
    .then(data => console.log('Updated product:', data))
    .catch(err => console.error('Error:', err));
}

// Delete a product (requires auth)
function testDeleteProduct(id) {
  fetch(`http://localhost:3000/products/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${authToken}`
    }
  })
    .then(res => res.json())
    .then(data => console.log('Deleted product:', data))
    .catch(err => console.error('Error:', err));
}

// GET requests don't require auth unless you enforce it
function testGetProducts() {
  fetch('http://localhost:3000/products')
    .then(res => res.json())
    .then(data => console.log('All products:', data))
    .catch(err => console.error('Error:', err));
}

function testGetProduct(id) {
  fetch(`http://localhost:3000/products/${id}`)
    .then(res => res.json())
    .then(data => console.log('Single product:', data))
    .catch(err => console.error('Error:', err));
}

