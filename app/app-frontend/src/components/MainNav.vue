<template>
    <header class="w-full text-sm">
        <div class="fixed top-0 left-0 w-full h-16 bg-brand-bg-1">
            <div class="flex flex-nowrap h-full border-b border-solid border-brand-gray-1 px-8 mx-auto">
                <a class="flex h-full items-center text-xl" :href="companyURL" target="_blank">{{ companyName }}</a>
                <nav class="ml-12 h-full">
                    <ul class="flex h-full list-none">
                        <li class="h-full ml-0">
                            <a class="flex h-full items-center py-2.5" href="#">Home</a>
                        </li>
                        <li class="h-full ml-9">
                            <a class="flex h-full items-center py-2.5" href="#">Books</a>
                        </li>
                        <li class="h-full ml-9">
                            <a class="flex h-full items-center py-2.5" href="#">About</a>
                        </li>
                    </ul>
                </nav>
                <div class="ml-auto flex h-full items-center">
                    <profile-image v-if="isLoggedIn"></profile-image>
                    <action-button v-else text="Sign in" @click="loginUser"></action-button>
                </div>
            </div>
        </div>
    </header>
</template>

<script>
import ActionButton from "@/components/ActionButton.vue";
import ProfileImage from "@/components/ProfileImage.vue";

export default {
    name: 'MainNav',
    components: {
        ActionButton,
        ProfileImage,
    },
    data() {
        return {
            companyName: "My Book Library",
            companyURL: "https://www.google.com",
            isLoggedIn: false,
        };
    },
    methods: {
        loginUser() {
            this.isLoggedIn = true;
        }
    }
}
</script>