<script lang="ts">
    import PlaylistHeader from "./PlaylistHeader.svelte";
    let sortBy = "artist";
    let sortDirection: "asc" | "desc" | null = null;

    function handleSort(field: string) {
        if (sortBy === field) {
            sortDirection = sortDirection === "asc" ? "desc" : "asc";
        } else {
            sortBy = field;
            sortDirection = "asc";
        }
    }
</script>

<div class="playlist-container">
    <table class="table table-xs table-zebra">
        <thead>
            <tr>
                <PlaylistHeader
                    sortable
                    sortDirection={sortBy === "artist" ? sortDirection : null}
                    onSort={() => handleSort("artist")}
                >
                    Artist / Album
                </PlaylistHeader>

                <PlaylistHeader
                    sortable
                    sortDirection={sortBy === "track" ? sortDirection : null}
                    onSort={() => handleSort("track")}
                >
                    Track
                </PlaylistHeader>

                <PlaylistHeader
                    sortable
                    sortDirection={sortBy === "title" ? sortDirection : null}
                    onSort={() => handleSort("title")}
                >
                    Title
                </PlaylistHeader>

                <PlaylistHeader
                    sortable
                    sortDirection={sortBy === "duration" ? sortDirection : null}
                    onSort={() => handleSort("duration")}
                >
                    Duration
                </PlaylistHeader>


            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Mastodon - Leviathan [Japan]</td>
                <td>01</td>
                <td>Blood and Thunder</td>
                <td>3:50</td>
            </tr>
            <!-- row 2 -->
            <tr>
                <td>Hart Hagerty</td>
                <td>Desktop Support Technician</td>
                <td>Purple</td>
                <td>-</td>
            </tr>
            <!-- row 3 -->
            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>

            <tr>
                <td>-</td>
                <td>-</td>
                <td>-</td>
                <td>-</td>
            </tr>
        </tbody>
    </table>
</div>

<style>
    .playlist-container {
        height: 100%;
        overflow-y: auto;
    }

    .sticky-header th {
        position: sticky;
        top: 0;
        z-index: 20;
    }
</style>
