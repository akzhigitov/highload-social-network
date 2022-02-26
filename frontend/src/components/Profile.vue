<template>

  <v-card>
    <v-parallax
        height="50"
        src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
    ></v-parallax>
    <v-card-text>
      <v-avatar size="88">
        <v-img v-if="user.sex===1" class="mb-4"
               src='https://avataaars.io/?avatarStyle=Circle&topType=LongHairBigHair&accessoriesType=Blank&hairColor=BrownDark&facialHairType=Blank&clotheType=GraphicShirt&clotheColor=White&graphicType=Deer&eyeType=Surprised&eyebrowType=AngryNatural&mouthType=Eating&skinColor=DarkBrown'></v-img>
        <v-img v-else class="mb-4"
               src='https://avataaars.io/?avatarStyle=Circle&topType=ShortHairTheCaesarSidePart&accessoriesType=Kurt&hairColor=BlondeGolden&facialHairType=BeardMedium&facialHairColor=Brown&clotheType=ShirtCrewNeck&clotheColor=Gray01&eyeType=Surprised&eyebrowType=RaisedExcited&mouthType=Eating&skinColor=DarkBrown'></v-img>
      </v-avatar>
      <h3 v-if="!isEdit" class="headline mb-2">
        {{ name }}
      </h3>
      <div v-else>
        <v-text-field
            v-model="user.firstName"
            label="FirstName"
        ></v-text-field>
        <v-text-field
            v-model="user.lastName"
            label="LastName"
        ></v-text-field>
      </div>
      <v-text-field
          v-model="user.city"
          :readonly="!isEdit"
          label="City"
      ></v-text-field>
      <v-menu
          ref="menu"
          v-model="menu"
          :close-on-content-click="false"
          :disabled="!isEdit"
          min-width="auto"
          offset-y
          transition="scale-transition">
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
              v-model="user.birthday"
              label="Birthday date"
              readonly
              v-bind="attrs"
              v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
            v-model="user.birthday"
            :active-picker.sync="activePicker"
            :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
            min="1950-01-01"
            @change="saveBD"
        ></v-date-picker>
      </v-menu>

      <v-select
          v-model="user.sex"
          :items="sexItems"
          :readonly="!isEdit"
          item-text="sex"
          item-value="id"
          label="Sex"
      ></v-select>
      <v-text-field
          v-model="user.interests"
          :readonly="!isEdit"
          label="Interests"
      ></v-text-field>

      <div v-if="!isMe">
        <v-btn
            v-if="isFollowing"
            class="ml-2 mt-5"
            outlined
            rounded
            small
            @click="unfollow">
          Unfollow
        </v-btn>
        <v-btn
            v-else
            class="ml-2 mt-5"
            outlined
            rounded
            small
            @click="follow">
          Follow
        </v-btn>
      </div>
      <div v-else>
        <v-btn v-if="!isEdit"
               class="ml-2 mt-5"
               outlined
               rounded
               small
               @click="edit">
          Edit
        </v-btn>
        <v-btn v-else
               class="ml-2 mt-5"
               outlined
               rounded
               small
               @click="save">
          Save
        </v-btn>
      </div>
    </v-card-text>
    <v-divider></v-divider>

    <v-card>
      <v-toolbar
          color="primary"
          dark
          flat
      >
        <v-toolbar-title>Friends</v-toolbar-title>
      </v-toolbar>
      <v-tabs vertical>
        <v-tab>
          Following
        </v-tab>
        <v-tab>
          Followers
        </v-tab>

        <v-tab-item>
          <v-card flat>
            <v-card-text>
              <user-tree-view :users="user.following"></user-tree-view>
            </v-card-text>
          </v-card>
        </v-tab-item>
        <v-tab-item>
          <v-card flat>
            <v-card-text>
              <user-tree-view :users="user.followers"></user-tree-view>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-card>
  </v-card>

</template>

<script>

import axios from "@/plugins/axios";

export default {
  data() {
    return {
      activePicker: null,
      menu: false,
      isFollowing: false,
      isEdit: false,
      sexItems: [
        {sex: 'Man', id: 0},
        {sex: 'Woman', id: 1},
      ],
    }
  },
  props: {
    user: Object
  },
  created() {
    this.isFollowing = this.calculateIsFollowing()
  },
  methods: {
    saveBD() {

    },
    edit() {
      this.isEdit = true
    },
    async save() {
      const token = this.$store.getters.token
      const config = {
        headers: {Authorization: `Bearer ${token}`}
      };

      console.log(this.user)
      await axios.put(`/users/${this.$props.user.id}`, this.user, config)
      this.isEdit = false
    },
    follow() {
      this.$emit('follow', this.user.id)
      this.isFollowing = true

    },
    unfollow() {
      this.$emit('unfollow', this.user.id)
      this.isFollowing = false
    },
    calculateIsFollowing() {
      const userID = this.$store.getters.userID
      if (!this.user.followers) {
        return false
      }
      for (let value of this.user.followers) {
        if (value.id == userID) {
          return true
        }
      }
      return false
    }
  },
  computed: {
    name() {
      return this.user.firstName + " " + this.user.lastName
    },
    isMe() {
      const userID = this.$store.getters.userID
      return this.user.id == userID
    },
  },
}
</script>

