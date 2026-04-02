<template>
  <div v-if="totalPages > 1" class="flex justify-center gap-2 mt-8">
    <button @click="$emit('change', page - 1)" :disabled="page <= 1" class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-30 hover:bg-gray-50">Prev</button>
    <span class="px-3 py-1 text-sm text-gray-500">{{ page }} / {{ totalPages }}</span>
    <button @click="$emit('change', page + 1)" :disabled="page >= totalPages" class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-30 hover:bg-gray-50">Next</button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  page: { type: Number, required: true },
  total: { type: Number, required: true },
  size: { type: Number, default: 10 },
})

defineEmits(['change'])

const totalPages = computed(() => Math.ceil(props.total / props.size))
</script>
