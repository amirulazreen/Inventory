<template>
  <div class="containerx">
    <h2 class="header">Inventory Table</h2>
    <div class="search_column">
      <input type="text" placeholder="search..." v-model="search" />
      <select v-model="selectedOption">
        <option value="ID">ID</option>
        <option value="Item">Item</option>
        <option value="Price">Price</option>
        <option value="Quantity">Quantity</option>
      </select>
      <RouterLink to="students/create" class="btn btn-primary">
        Add item
      </RouterLink>
    </div>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>No</th>
            <th @click="handleSort('ID')">ID</th>
            <th @click="handleSort('item')">Item</th>
            <th @click="handleSort('price')">Price</th>
            <th @click="handleSort('quantity')">Quantity</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in filteredInventory"
            :key="item.ID"
            @click="goToEdit(item.ID)"
          >
            <td>{{ index + 1 }}</td>
            <td>{{ item.ID }}</td>
            <td>{{ item.item }}</td>
            <td>{{ "RM " + item.price }}</td>
            <td>{{ item.quantity }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "inventory",
  data() {
    return {
      search: "",
      selectedOption: "Item",
      inventory: [],
    };
  },
  mounted() {
    this.getInventory();
  },
  methods: {
    getInventory() {
      axios
        .get("http://localhost:8080/inventory")
        .then((response) => {
          this.inventory = response.data;
        })
        .catch((error) => {
          console.error("Error fetching inventory:", error);
        });
    },
    deleteStudent(itemid) {
      if (confirm("Delete?")) {
        axios
          .delete(`http://localhost:8080/delete-inventory/${itemid}`)
          .then(() => {
            this.getInventory();
          });
      }
    },
    goToEdit(itemid) {
      this.$router.push({ path: `/students/${itemid}` });
    },
    handleSort(column) {
      if (this.sortBy === column) {
        this.sortDirection = this.sortDirection === "asc" ? "desc" : "asc";
      } else {
        this.sortBy = column;
        this.sortDirection = "asc";
      }

      this.inventory.sort((a, b) => {
        let aValue = a[column];
        let bValue = b[column];

        if (aValue < bValue) {
          return this.sortDirection === "asc" ? -1 : 1;
        } else if (aValue > bValue) {
          return this.sortDirection === "asc" ? 1 : -1;
        }
        return 0;
      });
    },
  },
  computed: {
    filteredInventory() {
      const searchTerm = this.search.toLowerCase();
      const selectedOption = this.selectedOption;

      return this.inventory.filter((item) => {
        switch (selectedOption) {
          case "Item":
            return item.item.toLowerCase().includes(searchTerm);
          case "Price":
            return item.price.toString().includes(searchTerm);
          case "ID":
            return item.ID.toString().includes(searchTerm);
          case "Quantity":
            return item.quantity.toString().includes(searchTerm);
          default:
            return true;
        }
      });
    },
  },
};
</script>
