_lang := `v=$(awk -F'"' '/^lang[[:space:]]*=/ {print $2; exit}' ~/.leetcode/leetcode.toml 2>/dev/null); echo "${v:-typescript}"`

# Open a LeetCode problem in your editor (right pane: problem, left pane: edit).
# Usage: just work <number> [language]
#        just work 42         # uses lang from ~/.leetcode/leetcode.toml
#        just work 42 rust    # override language
#        just work daily      # today's daily challenge (-d)
#        just work daily rust # override lang for daily
work number language=_lang:
    #!/usr/bin/env bash
    set -euo pipefail
    : "${SUPACODE_TAB_ID:?run from inside supacode}"
    : "${SUPACODE_SURFACE_ID:?}"
    if [ "{{number}}" = "daily" ]; then
        pick_arg="-d"
        edit_arg="-d -l {{language}}"
    else
        pick_arg="{{number}}"
        edit_arg="{{number}} -l {{language}}"
    fi
    supacode surface split -d h -i "clear && leetcode pick $pick_arg | cat"
    supacode surface focus -s "$SUPACODE_SURFACE_ID" -i "leetcode edit $edit_arg"
