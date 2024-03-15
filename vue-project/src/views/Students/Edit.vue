<template>
  <div class="container mt-5">
    <div class="card">
      <div class="card-header">
        <h4>Edit item</h4>
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
          <button type="button" @click="saveItem" class="btn btn-primary">
            Update
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "EditItem",
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
  mounted() {
    const itemId = this.$route.params.id;
    this.getItemData(itemId);
  },
  methods: {
    getItemData(itemId) {
      axios
        .get(`http://localhost:8080/update-inventory/${itemId}`)
        .then((response) => {
          this.model = response.data;
        })
        .catch((error) => {
          console.error("Error fetching item data:", error);
        });
    },
    saveItem() {
      axios
        .put(
          `http://localhost:8080/update-inventory/${this.$route.params.id}`,
          this.model
        )
        .then((response) => {
          console.log(response);
          alert("Item updated successfully");
          this.$router.push({
            path: `/students/${this.$route.params.id}`,
          });
        })
        .catch((error) => {
          console.error("Error updating item:", error);
        });
    },
  },
};
</script>
