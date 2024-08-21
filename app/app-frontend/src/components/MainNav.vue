<template>
    <header class="w-full text-sm">
        <div class="fixed top-0 left-0 w-full h-16 bg-brand-bg-1">
            <div class="flex flex-nowrap h-full border-b border-solid border-brand-gray-1 px-8 mx-auto">
                <a class="flex h-full items-center text-xl" :href="companyURL" target="_blank">{{ companyName }}</a>
<!--                <nav class="ml-12 h-full">-->
<!--                    <ul class="flex h-full list-none">-->
<!--                        <li class="first:ml-0 ml-9 h-full" v-for="menuItem in menuItems" v-bind:key="menuItem">-->
<!--                            <a href="#" class="flex h-full items-center py-2.5">{{ menuItem }}</a>-->
<!--                        </li>-->
<!--                    </ul>-->
<!--                </nav>-->
<!--                <div class="ml-auto flex h-full items-center">-->
<!--                    <profile-image v-if="isLoggedIn"></profile-image>-->
<!--                    <action-button v-else text="Sign in" @click="loginUser"></action-button>-->
<!--                </div>-->
            </div>
        </div>
    </header>
</template>

<script>
export default {
    name: 'MainNav',
    data() {
        return {
            companyName: "My Book Library",
            companyURL: "https://www.google.com",
        };
    },
}
</script>