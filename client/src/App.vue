<template>
  <Suspense>
    <File />
    <template #fallback> Loading ... </template>
  </Suspense>
</template>

<script setup>
import File from '@/components/File.vue'
</script>
