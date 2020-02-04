<template>
    <header>
        <div v-bind:key="tab.name" v-for="tab in tabs">
            <TopbarTab v-bind:tab="tab" v-bind:tabs="tabs"/>
        </div>
    </header>
</template>

<script>
import TopbarTab from './TopbarTab.vue'

export default {
    name: 'Topbar',
    components: {
        TopbarTab
    },
    props: ["tabs"]
}
</script>

<style scoped>
header {
    font-size: var(--header-font-size);
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    height: var(--header-height);
    background-color: var(--main-light-color);
    display: flex;
    color: white;
}
</style>
