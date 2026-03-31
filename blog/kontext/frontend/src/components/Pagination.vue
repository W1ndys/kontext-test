<template>
  <div v-if="totalPages > 1" class="flex items-center justify-center gap-2 mt-8">
    <!-- Previous -->
    <button
      :disabled="currentPage <= 1"
      class="px-3 py-2 text-sm rounded-lg border border-gray-300 text-gray-600 hover:bg-gray-100 transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
      @click="changePage(currentPage - 1)"
    >
      上一页
    </button>

    <!-- Page Numbers -->
    <template v-for="page in displayPages" :key="page">
      <span v-if="page === '...'" class="px-2 py-2 text-sm text-gray-400">...</span>
      <button
        v-else
        :class="[
          'px-3 py-2 text-sm rounded-lg border transition-colors',
          page === currentPage
            ? 'bg-blue-600 text-white border-blue-600'
            : 'border-gray-300 text-gray-600 hover:bg-gray-100'
        ]"
        @click="changePage(page)"
      >
        {{ page }}
      </button>
    </template>

    <!-- Next -->
    <button
      :disabled="currentPage >= totalPages"
      class="px-3 py-2 text-sm rounded-lg border border-gray-300 text-gray-600 hover:bg-gray-100 transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
      @click="changePage(currentPage + 1)"
    >
      下一页
    </button>

    <!-- Info -->
    <span class="ml-4 text-sm text-gray-400">
      共 {{ totalItems }} 条
    </span>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: {
    type: Number,
    required: true
  },
  totalPages: {
    type: Number,
    required: true
  },
  totalItems: {
    type: Number,
    default: 0
  }
})

const emit = defineEmits(['page-change'])

const displayPages = computed(() => {
  const total = props.totalPages
  const current = props.currentPage
  const pages = []

  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    pages.push(1)
    if (current > 3) {
      pages.push('...')
    }
    const start = Math.max(2, current - 1)
    const end = Math.min(total - 1, current + 1)
    for (let i = start; i <= end; i++) {
      pages.push(i)
    }
    if (current < total - 2) {
      pages.push('...')
    }
    pages.push(total)
  }

  return pages
})

function changePage(page) {
  if (page >= 1 && page <= props.totalPages && page !== props.currentPage) {
    emit('page-change', page)
  }
}
</script>
