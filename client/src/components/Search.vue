<template>
    <div id="search" v-bind:class="{ shown: isShown }">
        <input id="search-bar" type="text" placeholder="Search music...">
    </div>
</template>

<script>
export default {
    name: 'Search',
    props: {
        isShown: Boolean
    }
}
</script>

<style scoped>
#search {
    display: none;
}

#search-bar {
    display: block;
    margin: 50px auto;
    padding: 0 15px;
    border-radius: 10px;
    border: none;
    font-size: 18px;
    width: 80%;
    line-height: 50px;
    font-family: 'Roboto', Helvetica, Arial, sans-serif;
}

.shown {
    display: block !important;
}
</style>
