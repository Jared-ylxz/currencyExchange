<template>
    <div>
        <p>{{ count }}</p>
        <p>￥{{ amount }}</p>
        <ul>
            <li v-for="item in list" :key="item.name">{{ item.name }}</li>
        </ul>
        <button @click="count++">运气加一</button>
        <p>运气值：{{ luck }}</p>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue' // 给ref添加类型标注

const count = ref(0)

const amount = ref<number | string>('壹佰') // 泛型
type ListItem = {
    id: number
    name: string
}
const list = ref<ListItem[]>([])
list.value = [{
    id: 1,
    name: 'apple'
}]

const user = reactive({ // 自动推导字段的类型
    id: 1,
    name: "Jared"
})

const luck = computed(()=>{ // 能自动推导出 const luck: ComputedRef<number>
    return count.value ** 2
})
</script>

<style scoped>

</style>