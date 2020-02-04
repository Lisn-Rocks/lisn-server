<template>
    <a 
        v-bind:class="{ active: tab.isActive }"
        v-on:click="activate"
    >
        {{ tab.name }}
    </a>
</template>

<script>
export default {
    name: 'TopbarTab',
    props: ["tab", "tabs"],
    methods: {
        activate() {
            this.tabs.forEach(tab => {
                tab.isActive = false;
            });
            this.tab.isActive = true;
        }
    }
}
</script>

<style scoped>
a {
    display: block;
    line-height: calc( var(--header-height) - 3px );
    font-size: var(--header-font-size);
    padding: 0 var(--standard-horizontal-padding);
    text-align: center;
    cursor: pointer;
    font-weight: 400;
}
.active {
    font-weight: 500;
    border-bottom: 3px solid white;
    transition: all 0.3s;
}
</style>
