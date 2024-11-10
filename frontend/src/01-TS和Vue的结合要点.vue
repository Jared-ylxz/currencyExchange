<template>
    <div>
        <p>{{ count }}</p>
        <p>￥{{ amount }}</p>
        <ul>
            <li v-for="item in list" :key="item.name">{{ item.name }}</li>
        </ul>
        <button @click="count++">运气加一</button>
        <p>运气值：{{ luck }}</p>
        <ul>
            <li v-for="item in bigCurrency" :key="item.name">{{ item.name }} 的汇率是 {{ item.rate }}</li>
        </ul>
        <input type="text" placeholder="请输入内容" @change="inputChange">
        <button @click="Click">按钮内容</button>
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

type Currency = {
    id: number
    name: string
    rate: number
}

const Currency = reactive<Currency[]>([
    {id: 1, name: "USD", rate:1},
    {id: 2, name: "RMB", rate:7.1785},
    {id: 2, name: "MOP", rate:8.0075},
    {id: 2, name: "HKD", rate:7.7743},
    {id: 2, name: "EUR", rate:0.9329},
    {id: 2, name: "SGD", rate:1.32636},
    {id: 2, name: "GBP", rate:0.7740},
    {id: 2, name: "JPY", rate:152.629},
]);

const bigCurrency = computed(() => {
    return Currency.filter(el => el.rate < 1)
})

const inputChange = (e: Event) =>{ // 事件类型
    // console.log((e.target as HTMLInputElement).value) // 类型断言
    console.log((<HTMLInputElement>e.target).value) // 类型断言
}

const Click = (e: Event) =>{
    console.log((e.target as HTMLButtonElement).innerText) // 获取按钮内容
}
</script>

<style scoped>

</style>