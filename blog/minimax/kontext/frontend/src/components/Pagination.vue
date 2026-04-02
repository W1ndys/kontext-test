<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  current: number
  total: number
  pageSize?: number
}>()

const emit = defineEmits<{
  (e: 'change', page: number): void
}>()

const totalPages = computed(() => Math.ceil(props.total / (props.pageSize || 10)))

const visiblePages = computed(() => {
  const pages: (number | string)[] = []
  const current = props.current
  const total = totalPages.value

  if (total <= 7) {
    for (let i = 1; i <= total; i++) pages.push(i)
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }
  return pages
})

function goTo(page: number) {
  if (page >= 1 && page <= totalPages.value && page !== props.current) {
    emit('change', page)
  }
}
</script>

<template>
  <div class="flex justify-center items-center gap-2 mt-8">
    <button
      @click="goTo(current - 1)"
      :disabled="current === 1"
      class="px-3 py-1 rounded border disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100"
    >
      上一页
    </button>
    <template v-for="page in visiblePages" :key="page">
      <span v-if="page === '...'" class="px-2">...</span>
      <button
        v-else
        @click="goTo(page as number)"
        :class="['px-3 py-1 rounded border', page === current ? 'bg-primary-600 text-white border-primary-600' : 'hover:bg-gray-100']"
      >
        {{ page }}
      </button>
    </template>
    <button
      @click="goTo(current + 1)"
      :disabled="current === totalPages"
      class="px-3 py-1 rounded border disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100"
    >
      下一页
    </button>
  </div>
</template>
