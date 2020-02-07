<template>
    <a 
        v-bind:class="{ active: isActive }"
        v-on:click="$emit('switch')"
    >
        {{ tab.name }}
    </a>
</template>

<script>
export default {
    name: 'TopbarTab',
    props: {
        tab: Object,
        isActive: Boolean
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
    transition: all 0.2s;
}
</style>
