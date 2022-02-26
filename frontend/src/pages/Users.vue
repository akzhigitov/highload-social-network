<template>
  <v-container fluid grid-list-xl>
    <v-data-table
        :headers="headers"
        :items="users"
        class="table"
        @click:row="open">
      <template v-slot:item.sex="{ item }" >
        {{ item.sex === 0 ? "Man":"Woman" }}
      </template>
    </v-data-table>
  </v-container>
</template>
<script>
import axios from "@/plugins/axios";

export default {

  data() {
    return {
      users: [],
      footerProps: {'items-per-page-options': [15, 30, 50, 100]},
      headers: [
        {
          text: 'First Name',
          value: 'firstName',
          align: 'left',
          sortable: true
        },
        {
          text: 'Last Name',
          value: 'lastName',
          align: 'left',
          sortable: true
        },
        {
          text: 'Birthday',
          value: 'birthday',
          align: 'left',
          sortable: true
        },
        {
          text: 'Sex',
          value: 'sex',
          align: 'left',
          sortable: true
        },
        {
          text: 'City',
          value: 'city',
          align: 'left',
          sortable: true
        },
      ]
    }
  },

  methods: {
    async fetchUsers() {
      try {
        const response = await axios.get('/users/')
        this.users = response.data
      } catch (e) {
        console.log(e)
      }
    },
    open(row) {
      this.$router.push(`/users/${row.id}`);
    }
  },
  mounted() {
    this.fetchUsers()
  }
}
</script>

<style>
.table {
  border-radius: 3px;
  background-clip: border-box;
  border: 1px solid rgba(0, 0, 0, 0.125);
  box-shadow: 1px 1px 1px 1px rgba(0, 0, 0, 0.21);
  background-color: transparent;
}
</style>
