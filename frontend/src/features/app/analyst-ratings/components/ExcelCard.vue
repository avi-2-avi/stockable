<template>
    <Card :cardClass="'cursor-pointer hover:bg-black/5 transition ' + cardClass" @click="exportToExcel">
        <div class="flex items-center space-x-4">
            <div class="p-2.5 bg-stock-500 rounded-full">
                <Sheet class="w-6 h-6 text-base" />
            </div>
            <div>
                <p class="text-lg font-medium font-outfit">Download to Excel</p>
                <p class="text-sm text-foreground/80">Click to download the data in Excel format.</p>
            </div>
        </div>
    </Card>
</template>

<script setup lang="ts">
import Card from '@/components/ui/Card.vue'
import { Sheet } from 'lucide-vue-next'
import * as XLSX from 'xlsx';

const props = defineProps({
    cardClass: {
        type: String,
        default: ''
    },
    fileName: {
        type: String,
        default: 'default.xlsx'
    },
    tableData: {
        type: Array,
        default: () => []
    }
});

const exportToExcel = () => {
    if (!props.tableData || props.tableData.length === 0) {
        return;
    }

    const ws = XLSX.utils.json_to_sheet(props.tableData);
    const wb = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(wb, ws, props.fileName);
    XLSX.writeFile(wb, props.fileName + '.xlsx');
}

</script>