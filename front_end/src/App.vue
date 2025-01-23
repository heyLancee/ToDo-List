<script setup>
import { ref } from 'vue'
import axios from 'axios'

const value = ref("")
const item_map = ref({})

get_list()

async function get_list() {
  const res = await axios({
    url: 'http://localhost:8080/api/get_list',
    method: 'get'
  })

  item_map.value = res.data
}


async function add() {
  await axios({
    url: 'http://localhost:8080/api/add_list',
    method: 'post',
    data: {
      value: value.value,
      isCompleted: false
    }
  })

  value.value = ''
  get_list()
}

async function update(index) {
  await axios({
    url: 'http://localhost:8080/api/update_list',
    method: 'post',
    data: {
      id: index
    }
  })
}

async function delete_item(index) {
  console.log(index)
  await axios({
    url: 'http://localhost:8080/api/delete_list',
    method: 'post',
    data: {
      id: index
    }
  })

  get_list()
}
</script>
<template>
  <div class="todo-app">
    <div class="title">Todo App</div>

    <div class="todo-form">
      <input
        v-model="value"
        type="text"
        class="todo-input"
        placeholder="Add a todo"
      />
      <div @click="add" class="todo-button">Add Todo</div>
    </div>

    <div
      v-for="(value, key) in item_map"
      :class="[value.isCompleted ? 'completed' : 'item']"
    >
      <div>
        <input @click="update(key)" v-model="value.isCompleted" type="checkbox" />
        <span class="name">{{ value.value }}</span>
      </div>

      <div @click="delete_item(key)" class="del">del</div>
    </div>
  </div>
</template>

<style scoped>
.todo-app {
  box-sizing: border-box;
  margin-top: 40px;
  margin-left: 1%;
  padding-top: 30px;
  width: 98%;
  height: 500px;
  background: #ffffff;
  border-radius: 5px;
}

.title {
  text-align: center;
  font-size: 30px;
  font-weight: 700;
}

.todo-form {
  display: flex;
  margin: 20px 0 30px 20px;
}

.todo-button {
  width: 100px;
  height: 52px;
  border-radius: 0 20px 20px 0;

  text-align: center;
  background: linear-gradient(
    to right,
    rgb(113, 65, 168),
    rgba(44, 114, 251, 1)
  );
  color: #fff;
  line-height: 52px;
  cursor: pointer;
  font-size: 14px;
  user-select: none;
}

.todo-input {
  padding: 0px 15px 0px 15px;
  border-radius: 20px 0 0 20px;
  border: 1px solid #dfe1e5;
  outline: none;
  width: 60%;
  height: 50px;
}

.item {
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 80%;
  height: 50px;
  margin: 8px auto;
  padding: 16px;
  border-radius: 20px;
  box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 20px;
}

.del {
  color: red;
  user-select: none;
  cursor: pointer;
}

.completed {
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 80%;
  height: 50px;
  margin: 8px auto;
  padding: 16px;
  border-radius: 20px;
  box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 20px;
  text-decoration: line-through;
  opacity: 0.4;
}
</style>
