<template>
    <div class="progress-bar">
        <div class="progress" v-bind:style="{ width: progress }"/>
    </div>
</template>

<script>
export default {
    name: 'ProgressBarMin',

    props: {
        currentSong: HTMLAudioElement
    },

    created() {
        this.currentSong.addEventListener(
            'timeupdate', () => this.$emit('upd')
        );
    },
    
    computed: {
        progress: function() {
            let cs = this.currentSong;
            return cs.currentTime / cs.duration * 100 + '%';
        }
    },
}
</script>

<style scoped>
.progress-bar {
    width: 100%;
    height: 3px;
    background-color: var(--secondary-white);
}

.progress-bar .progress {
    height: 100%;
    background-color: var(--accent-color);
}
</style>