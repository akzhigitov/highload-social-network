<template>
  <v-app id="login" class="secondary">
    <v-main>
      <v-container fill-height fluid>
        <v-layout align-center justify-center>
          <v-flex lg4 md4 sm8 xs12>
            <v-card class="elevation-12">
              <v-card-text>
                <div class="layout column align-center">
                  <h1 class="flex my-4 primary--text">HighLoad Social Network</h1>
                </div>
                <v-form>
                  <v-text-field
                      v-model="userEmail"
                      :error="error"
                      :rules="[rules.required]"
                      append-icon="person"
                      label="Login"
                      name="login"
                      type="text"/>
                  <v-text-field
                      id="password"
                      v-model="password"
                      :append-icon="hidePassword ? 'visibility_off' : 'visibility'"
                      :error="error"
                      :rules="[rules.required]"
                      :type="hidePassword ? 'password' : 'text'"
                      label="Password"
                      name="password"
                      @click:append="hidePassword = !hidePassword"/>
                </v-form>
              </v-card-text>

              <v-card-actions>

                <v-btn :loading="loading" color="secondary" @click="$router.push({ name: 'Registration' });">
                  Registration
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn :loading="loading" color="primary" @click="submit">Login</v-btn>
              </v-card-actions>

            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
      <v-snackbar
          v-model="showResult"
          :timeout="2000"
          top>
        {{ result }}
      </v-snackbar>
    </v-main>
  </v-app>
</template>

<script>
import {mapActions} from "vuex";

export default {

  data() {
    return {
      loading: false,
      userEmail: '',
      password: '',
      hidePassword: true,
      error: false,
      showResult: false,
      result: '',
      rules: {
        required: value => !!value || 'Required.'
      }
    }
  },

  methods: {
    ...mapActions(["LogIn"]),
    async submit() {

      const vm = this;
      if (!vm.userEmail || !vm.password) {

        vm.result = "Email and Password can't be null.";
        vm.showResult = true;

        return;
      }

      const User = {email: vm.userEmail, password: vm.password}

      try {
        await this.LogIn(User);
        await this.$router.push({name: 'Home'});
        this.showError = false
      } catch (error) {
        vm.error = true;
        vm.result = "Email or Password is incorrect.";
        vm.showResult = true;
        console.log(error)
      }

    }
  }
}
</script>

<style lang="css" scoped>
#login {
  height: 50%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
  content: "";
  z-index: 0;
}
</style>
