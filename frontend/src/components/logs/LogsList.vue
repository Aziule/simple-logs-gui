<template>
    <section>
        <table class="striped highlight grey lighten-5">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Time</th>
                    <th>
                        Level
                        <Level />
                    </th>
                    <th>
                        Message
                        <Message />
                    </th>
                    <th>Contextual Info</th>
                </tr>
            </thead>
            <tbody v-if="logs.length > 0">
                <LogItem v-for="(log, index) in filteredLogs" v-bind:log="log" v-bind:key="index"/>
            </tbody>
        </table>
        <div class="no-logs" v-if="logs.length === 0">
            No new logs at the moment
        </div>
    </section>
</template>

<script>
import { mapGetters } from 'vuex'
import LogItem from '@/components/logs/LogItem'
import Level from '@/components/logs/filters/Level'
import Message from '@/components/logs/filters/Message'

export default {
    name: 'LogsList',
    computed: {
        ...mapGetters([
            'logs',
            'filters'
        ]),
        filteredLogs () {
            return this.logs.filter((log) => {
                if (this.filters.level && this.filters.level !== log.level) {
                    return false
                }

                var content = log.message || log.raw

                if (this.filters.message) {
                    return content.toLowerCase().includes(this.filters.message.toLowerCase())
                }

                return true
            })
        }
    },
    components: {
        LogItem,
        Level,
        Message
    }
}
</script>

<style>
table th {
    vertical-align: top;
    font-weight: 500;
}
tbody tr:hover {
    background-color: #eee!important
}
.no-logs {
    position: absolute;
    top: 98px;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #eee;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: -1;
}
</style>
