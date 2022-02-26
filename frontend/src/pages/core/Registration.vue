<template>
  <v-app id="registration" class="secondary">
    <v-main>
      <v-main fill-height fluid>
        <v-layout align-center justify-center>
          <v-form
              ref="form"
              v-model="valid"
              lazy-validation
          >
            <v-text-field
                v-model="firstName"
                :counter="10"
                :rules="nameRules"
                label="First Name"
                required
            ></v-text-field>

            <v-text-field
                v-model="lastName"
                :counter="10"
                :rules="nameRules"
                label="Last Name"
                required
            ></v-text-field>

            <v-text-field
                v-model="email"
                :rules="emailRules"
                label="E-mail"
                required
            ></v-text-field>

            <v-text-field
                v-model="password"
                label="Password"
                required
                type='password'
            ></v-text-field>


            <v-menu
                ref="menu"
                v-model="menu"
                :close-on-content-click="false"
                min-width="auto"
                offset-y
                transition="scale-transition"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                    v-model="birthday"
                    label="Birthday date"
                    readonly
                    v-bind="attrs"
                    v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker
                  v-model="birthday"
                  :active-picker.sync="activePicker"
                  :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
                  min="1950-01-01"
                  @change="save"
              ></v-date-picker>
            </v-menu>

            <v-text-field
                v-model="city"
                label="City"
                required
            ></v-text-field>

            <v-text-field
                v-model="interests"
                label="Interests"
                required
            ></v-text-field>

            <v-select
                v-model="sex"
                :items="sexItems"
                :rules="[v => !!v || 'Sex is required']"
                label="Sex"
                required
            ></v-select>

            <v-btn @click="registration">
              Register
            </v-btn>
          </v-form>
        </v-layout>
      </v-main>
    </v-main>
  </v-app>
</template>

<script>
import axios from "@/plugins/axios";

export default {
  data: () => ({
    activePicker: null,
    menu: false,
    valid: true,
    firstName: '',
    lastName: '',
    nameRules: [
      v => !!v || 'Name is required',
      v => (v && v.length <= 10) || 'Name must be less than 10 characters',
    ],
    birthday: null,
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    password: '',
    sex: null,
    sexItems: [
      'Man',
      'Woman'
    ],
    city: '',
    interests: ''
  }),
  watch: {
    menu(val) {
      val && setTimeout(() => (this.activePicker = 'YEAR'))
    },
  },
  methods: {
    save(date) {
      this.$refs.menu.save(date)
    },
    async registration() {
      await axios.post("/auth/sign-up", {
        email: this.email,
        password: this.password,
        firstName: this.firstName,
        lastName: this.lastName,
        birthday: this.birthday,
        sex: this.sex === "Man" ? 0 : 1,
        city: this.city,
        interests: this.interests
      })

      await this.$router.push("/")
    },
  },
}
</script>

<style lang="css" scoped>
#registration {
  height: 50%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
  content: "";
  z-index: 0;
}
</style>
