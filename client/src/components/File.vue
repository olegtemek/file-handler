<template>
  <div class="container">
    <h1>FileHandlerAdmin</h1>
    <div class="file">
      <div class="file__tags">
        <div class="file__tag" v-for="(tag, index) of tags" :key="index">
          <button @click="getFilesByTag(tag)" :class="activeTag == tag ? 'active' : ''">
            {{ tag }}
          </button>
        </div>
      </div>

      <div class="file__wrapper">
        <div class="file__wrapper-item" v-for="(file, index) of files" :key="file.id">
          <span>{{ new Date(file.timestamp).toLocaleString() }}</span>
          <p>
            {{ file.text }}
          </p>
          <button @click="getText(index)">See log</button>
          <button @click="remove(file.id)" class="delete">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { UseFile } from '@/composables/file'
import { onMounted, ref } from 'vue'

const useFile = UseFile()

const tags = ref()
const files = ref()
const activeTag = ref()

onMounted(async () => {
  tags.value = await useFile.getTags()
})

const getFilesByTag = async (tag) => {
  files.value = await useFile.getFilesByTag(tag)
  activeTag.value = tag
  console.log(files.value)
}

const getText = async (index) => {
  const file = files.value[index]
  if (file?.text) {
    delete file.text
    return
  }
  const text = await useFile.getText(file.filepath)
  files.value[index].text = text
}

const remove = async (id) => {
  const result = await useFile.remove(id)
  if (result) {
    files.value = files.value.filter((item) => item.id != id)
    return
  }
  alert('ERROR!')
}
</script>
