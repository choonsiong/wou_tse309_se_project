<template>
  <section class="bg-green-100">
    <h2 class="pt-10 text-4xl mb-6 font-bold text-center">Users Administration</h2>
    <p class="max-w-xs mx-auto text-center text-gray-800 md:max-w-md">
      Note: All username in red are <strong>Administrator</strong> and user must be in <strong>active</strong> status in
      order to login.
    </p>
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <div v-for="(user, index) in users" :key="user.id">
        <div class="flex flex-col items-center justify-between w-full p-6 bg-green-500/30 rounded-lg md:flex-row">
          <p class="font-bold text-center text-veryDarkViolet md:text-left mr-2">#{{ index + 1 }}</p>
          <img class="h-14 w-14 mr-3 rounded-full p-[1.5px]  hover:scale-110 transition-transform duration-200 ease-out"
               :src="'https://i.pravatar.cc/150?img='+user.id" alt="profile image">
          <p v-if="user.is_admin" class="font-bold text-center text-rose-500 md:text-left">{{ user.first_name }}
            {{ user.last_name }}</p>
          <p v-else class="font-bold text-center text-veryDarkViolet md:text-left">{{ user.first_name }}
            {{ user.last_name }}</p>&nbsp;
          <div class="flex flex-col items-center justify-end flex-1 space-x-4 space-y-2 md:flex-row md:space-y-0">
            <div class="font-bold text-gray-800 mr-4">{{ user.email }}</div>
            <button v-if="user.active"
                    class="p-2 px-8 text-white bg-blue-500 rounded-lg hover:opacity-70 focus:outline-none">Active
            </button>
            <button v-else class="p-2 px-8 text-white bg-gray-500 rounded-lg hover:opacity-70 focus:outline-none">
              Inactive
            </button>
            <button class="p-2 px-8 text-white bg-cyan rounded-lg hover:opacity-70 focus:outline-none"
                    @click="handleEditUser(user.id)">Edit
            </button>
            <button class="p-2 px-8 text-white bg-red rounded-lg hover:opacity-70 focus:outline-none">Delete</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="isAddUser" class="max-w-2xl mx-auto p-6 space-y-6">
      <h2 class="pt-10 text-4xl mb-6 font-bold text-center">Add New User</h2>
      <p class="max-w-xs mx-auto text-center text-gray-800 md:max-w-md">
        Note: All fields are required.
      </p>
      <div>
        <form class="flex flex-col items-center justify-between w-full p-6 bg-green-500/30 rounded-lg md:flex-col">
          <input type="text" placeholder="First Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input type="text" placeholder="Last Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input type="email" placeholder="Email"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input type="password" placeholder="Password"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input type="password" placeholder="Confirm Password"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <div class="px-6 pt-3 pb-2 mb-4">
            <input type="checkbox"> User Active&nbsp;&nbsp;
            <input type="checkbox"> Administrator
          </div>
          <button class="w-full p-3 px-8 text-white bg-red-600 rounded-lg hover:opacity-70 focus:outline-none">Submit
          </button>
        </form>
      </div>
    </div>
    <div v-if="isEditUser" class="max-w-2xl mx-auto p-6 space-y-6">
      <h2 class="pt-10 text-4xl mb-6 font-bold text-center">Edit User ID ({{ editUserID }})</h2>
      <p class="max-w-xs mx-auto text-center text-gray-800 md:max-w-md">
        Note: All fields are required.
      </p>
      <div>
        <form @submit.prevent="submitEditForm"
              class="flex flex-col items-center justify-between w-full p-6 bg-green-500/30 rounded-lg md:flex-col">
          <input v-model="editUserFirstName" type="text" placeholder="First Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input v-model="editUserLastName" type="text" placeholder="Last Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input v-model="editUserEmail" type="email" placeholder="Email"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <div class="w-full flex-1">
            <p class="mb-2">Leave the password field empty to use existing password.</p>
            <input v-model="editUserPassword" type="password" placeholder="Password"
                   class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          </div>
          <div class="w-full flex-1">
            <p class="mb-2">Leave the confirm password field empty to use existing password.</p>
            <input v-model="editUserConfirmPassword" type="password" placeholder="Confirm Password"
                   class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          </div>
          <div class="px-6 pt-3 pb-2 mb-4">
            <input type="checkbox" :checked="editUserIsActive" @change="setEditUserIsActive"> User Active&nbsp;&nbsp;
            <input type="checkbox" :checked="editUserIsAdmin" @change="setEditUserIsAdmin"> Administrator
          </div>
          <div class="flex flex-row justify-between w-full gap-4">
            <button class="w-full p-3 px-8 text-white bg-blue-600 rounded-lg hover:opacity-70 focus:outline-none"
                    @click="cancelEdit">Cancel
            </button>
            <button class="w-full p-3 px-8 text-white bg-red rounded-lg hover:opacity-70 focus:outline-none">Submit
            </button>
          </div>
        </form>
      </div>
    </div>
    <div class="p-10">&nbsp;</div>
  </section>
</template>

<script>
import Security from '@/security.js'
import appEnvironment from '@/environment.js'
import notie from 'notie'

export default {
  name: 'AllUsers',
  data() {
    return {
      users: [],
      isAddUser: false,
      isEditUser: false,
      editUserID: '',
      editUserFirstName: '',
      editUserLastName: '',
      editUserEmail: '',
      editUserPassword: '',
      editUserConfirmPassword: '',
      editUserIsActive: 0,
      editUserIsAdmin: 0
    }
  },
  methods: {
    handleEditUser(userId) {
      notie.alert({
        type: 'info',
        text: 'Editing user id: ' + userId
      })
      this.isEditUser = true
      this.editUserID = userId
      this.editUserFirstName = this.users.filter((user) => user.id === userId)[0].first_name
      this.editUserLastName = this.users.filter((user) => user.id === userId)[0].last_name
      this.editUserEmail = this.users.filter((user) => user.id === userId)[0].email
      this.editUserPassword = ''
      this.editUserConfirmPassword = ''
      //this.editUserPassword = 'this.users.filter((user) => user.id === userId)[0].password'
      //this.editUserConfirmPassword = this.users.filter((user) => user.id === userId)[0].password
      console.log(this.users.filter((user) => user.id === userId)[0].active)
      this.editUserIsActive = this.users.filter((user) => user.id === userId)[0].active
      console.log(this.users.filter((user) => user.id === userId)[0].is_admin)
      this.editUserIsAdmin = this.users.filter((user) => user.id === userId)[0].is_admin
    },
    cancelEdit() {
      this.isEditUser = false
    },
    setEditUserIsActive() {
      if (this.editUserIsActive === 0) {
        this.editUserIsActive = 1
      } else if (this.editUserIsActive === 1) {
        this.editUserIsActive = 0
      }
    },
    setEditUserIsAdmin() {
      if (this.editUserIsAdmin === 0) {
        this.editUserIsAdmin = 1
      } else if (this.editUserIsAdmin === 1) {
        this.editUserIsAdmin = 0
      }
    },
    submitEditForm() {
      console.log('submitEditForm')

      this.isEditUser = false

      console.log(this.editUserFirstName)
      console.log(this.editUserLastName)
      console.log(this.editUserEmail)
      console.log(this.editUserPassword)
      console.log(this.editUserConfirmPassword)
      console.log(this.editUserIsActive)
      console.log(this.editUserIsAdmin)

      if (this.editUserPassword !== this.editUserConfirmPassword) {
        notie.alert({
          type: 'error',
          text: 'password entered does not match',
        })
        return
      }

      const payload = {
        id: parseInt(String(this.editUserID), 10),
        first_name: this.editUserFirstName,
        last_name: this.editUserLastName,
        email: this.editUserEmail,
        password: this.editUserPassword,
        active: this.editUserIsActive,
        is_admin: this.editUserIsAdmin,
      }

      console.log(payload)

      fetch(appEnvironment.apiURL() + '/admin/users/save', Security.requestOptions())
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            notie.alert({
              type: 'error',
              text: jsonResp.message
            })
          } else {
            notie.alert({
              type: 'success',
              text: 'user saved successfully',
            })
          }
        })
        .catch((err) => {
          notie.alert({
            type: 'error',
            text: err
          })
        })
    }
  },
  beforeMount() {
    Security.requireAdmin()

    fetch(appEnvironment.apiURL() + '/admin/users/all', Security.requestOptions(''))
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          //console.log('error: ' + jsonResp.message)
          notie.alert({
            type: 'error',
            text: jsonResp.message
          })
        } else {
          //console.log(jsonResp)
          this.users = jsonResp.data.users
        }
      })
      .catch((err) => {
        notie.alert({
          type: 'error',
          text: err
        })
      })
  }
}
</script>

<style scoped>

</style>