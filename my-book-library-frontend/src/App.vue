<template>
    <main-header></main-header>
    <main-content></main-content>
    <main-footer></main-footer>
</template>

<script>
import MainHeader from "@/components/MainHeader.vue";
import MainFooter from "@/components/MainFooter.vue";
import MainContent from "@/components/MainContent.vue";

export default {
    name: 'App',
    components: {
        MainHeader,
        MainContent,
        MainFooter,
    }
}
</script>
