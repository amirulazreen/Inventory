<template>
  <div class="container mt-5">
    <div class="card">
      <div class="card-header">
        <h4>Item detail</h4>
      </div>
      <div v-if="loading">Loading...</div>
      <div v-else>
        <div class="card-body">
          <div class="mb-3">
            <h2>ID: {{ product.ID }}</h2>
          </div>
          <div class="mb-3">
            <h2>Item: {{ product.item }}</h2>
          </div>
          <div class="mb-3">
            <h2>Price: {{ product.price }}</h2>
          </div>
          <div class="mb-3">
            <h2>Quantity: {{ product.quantity }}</h2>
          </div>
          <div class="mb-3">
            <h2>Supplier Name: {{ product.supplier.name }}</h2>
          </div>
          <div class="mb-3">
            <h2>Address: {{ product.supplier.address }}</h2>
          </div>
          <div class="mb-3">
            <h2>Contact: {{ product.supplier.tel }}</h2>
          </div>
          <div class="mb-3">
            <router-link
              :to="{ path: '/students/' + product.ID + '/edit' }"
              class="btn btn-success"
            >
              Edit
            </router-link>
            <button
              type="button"
              @click="deleteStudent(product.ID)"
              class="btn btn-danger"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      product: null,
      loading: true,
    };
  },
  mounted() {
    this.fetchProduct();
  },
  methods: {
    fetchProduct() {
      const id = this.$route.params.id;
      fetch(`http://localhost:8080/inventory/${id}`)
        .then((response) => response.json())
        .then((data) => {
          this.product = data;
          this.loading = false;
        })
        .catch((error) => {
          console.error("Error fetching product:", error);
          this.loading = false;
        });
    },
    deleteStudent(itemid) {
      if (confirm("Delete?")) {
        axios
          .delete(`http://localhost:8080/delete-inventory/${itemid}`)
          .then(() => {
            this.fetchProduct();
            this.$router.push({ path: "/students" });
          });
      }
    },
    goToEdit(itemid) {
      this.$router.push({ path: `/students/${itemid}` });
    },
  },
};
</script>
