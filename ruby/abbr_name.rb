def abbrev_name(name)
    abbr = name.split(' ')
    return abbr[0][0].upcase + '.' + abbr[1][0].upcase
end