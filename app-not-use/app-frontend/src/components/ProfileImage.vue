<template>
    <img :src="imageLink" class="h-8 w-8 rounded-3xl object-contain" alt="User profile image"/>
</template>

<script>
export default {
    name: 'ProfileImage',
    data() {
        return {
            imageLink: "https://mighty.tools/mockmind-api/content/human/49.jpg",
        }
    }
}
</script>