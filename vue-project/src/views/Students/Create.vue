<template>
  <div class="container mt-5">
    <div class="card">
      <div class="card-header">
        <h4>Add item</h4>
      </div>
      <div class="card-body">
        <div class="mb-3">
          <label for="">Item</label>
          <input type="text" v-model="model.item" class="form-control" />
        </div>
        <div class="mb-3">
          <label for="">Price</label>
          <input type="number" v-model="model.price" class="form-control" />
        </div>
        <div class="mb-3">
          <label for="">Quantity</label>
          <input type="number" v-model="model.quantity" class="form-control" />
        </div>
        <div class="mb-3">
          <label for="">Supplier Name</label>
          <input
            type="text"
            v-model="model.supplier.name"
            class="form-control"
          />
        </div>
        <div class="mb-3">
          <label for="">Supplier Contact</label>
          <input
            type="text"
            v-model="model.supplier.tel"
            class="form-control"
          />
        </div>
        <div class="mb-3">
          <label for="">Supplier Address</label>
          <input
            type="text"
            v-model="model.supplier.address"
            class="form-control"
          />
        </div>
        <div class="mb-3">
          <button type="button" @click="saveStudent" class="btn btn-primary">
            Save
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "StudentsCreate",
  data() {
    return {
      model: {
        item: "",
        price: 0,
        quantity: 0,
        supplier: {
          name: "",
          address: "",
          tel: "",
        },
      },
    };
  },
  methods: {
    saveStudent() {
      axios
        .post("http://localhost:8080/add-inventory", this.model)
        .then((res) => {
          console.log(res);
          alert("Item added");
          this.$router.push({ path: "/students" });
        })
        .catch(function (error) {
          if (error.response) {
            console.log(error.response.data);
            console.log(error.response.status);
            console.log(error.response.headers);
          } else if (error.request) {
            console.log(error.request);
          } else {
            console.log("Error", error.message);
          }
          console.log(error.config);
        });
    },
  },
};
</script>
