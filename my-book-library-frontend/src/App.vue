<template>
    <main-header></main-header>
    <div>
        <router-view />
    </div>
    <main-footer></main-footer>
</template>

<script>
import MainHeader from "@/components/MainHeader.vue";
import MainFooter from "@/components/MainFooter.vue";

export default {
    name: 'App',
    components: {
        MainHeader,
        MainFooter,
    }
}
</script>
