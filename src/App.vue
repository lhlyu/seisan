<template>
    <main>
        <textarea v-model='txt' rows='10' cols='100'>

        </textarea>
        <br>
        <button @click='click'>执行</button>
        <br>
        <br>
        <textarea v-model='item' rows='20' cols='100'>
        </textarea>
    </main>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const txt = ref<string>('')
const item = ref('')

const click = () => {
    const set = new Set()
    txt.value.split('').map((v:string) => {
        const val = v.trim()
        if (val.length) {
            set.add(val)
        }
    })
    item.value = [...set].join('\n')
}
</script>

<style>
main {
    padding: 50px;
}
</style>